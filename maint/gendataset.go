package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	os.Exit(_main())
}

func _main() (st int) {
	defer func() {
		if err := recover(); err != nil {
			st = 1
		}
	}()

	var err error

	fmt.Println(" + checking for 'woothee'...")
	if _, err := os.Stat("woothee"); err != nil {
		if err = exec.Command("git", "clone", "git://github.com/woothee/woothee.git").Run(); err != nil {
			fmt.Println(err)
			return 1
		}
	}

	fmt.Println(" + writing dataset...")
	if err = writeDataset(); err != nil {
		fmt.Printf("Failed to write dataset: %s\n", err)
		return 1
	}

	fmt.Println(" + writing testset...")
	if err = writeTestset(); err != nil {
		fmt.Printf("Failed to generate testset: %s\n", err)
		return 1
	}

	fmt.Println(" + formatting output...")
	if err = exec.Command("go", "fmt", ".").Run(); err != nil {
		fmt.Printf("Failed to run go cmd: %s\n", err)
		return 1
	}

	fmt.Println("*** all done ***")
	return 0
}

func writeDataset() error {
	src, err := os.Open("woothee/dataset.yaml")
	if err != nil {
		return err
	}
	defer src.Close()

	fi, err := src.Stat()
	if err != nil {
		return err
	}

	dst, err := os.OpenFile("dataset.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()

	buf := make([]byte, fi.Size())
	if _, err := io.ReadFull(src, buf); err != nil {
		return err
	}

	var data []struct {
		Label    string `yaml:"label"`
		Name     string `yaml:"name"`
		Category string `yaml:"category"`
		Type     string `yaml:"type"`
		Vendor   string `yaml:"vendor,omitempty"`
		Os       string `yaml:"os,omitempty"`
	}

	if err = yaml.Unmarshal(buf, &data); err != nil {
		return err
	}

	t, err := template.New("dataset").Parse(`/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

var DefaultDataSet DataSet = DataSet {
{{range .}}	"{{.Label}}": &Result {
		Name: "{{.Name}}",
		Type: "{{.Category}}",
{{if eq .Type "browser"}}		Vendor: "{{.Vendor}}",
{{else if eq .Type "os"}}		Category: "{{.Category}}",
{{else if eq .Type "full"}}		Vendor: "{{.Vendor}}",
		Category: "{{.Category}}",
		Os: {{with .Os}}"{{.}}"{{else}}"UNKNOWN"{{end}},
{{end}} },
{{end}}
}
`)
	if err != nil {
		return err
	}

	if err = t.Execute(dst, data); err != nil {
		return err
	}

	return nil
}

func writeTestset() error {
	tmpl := `/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)
func Test_{{.TestName}}(t *testing.T) {
	var result *Result
	var err error
{{range .Tests}} result, err = Parse(` + "`{{.Target}}`" + `)
	if err != nil {
		t.Errorf(` + "`Failed to parse '{{.Target}}': %s`" + `, err)
	} else {
`

	for _, field := range []string{"Category", "Name", "Os", "OsVersion", "Version"} {
		tmpl += fmt.Sprintf(`		if result.%s != "{{.%s}}" {
			t.Errorf("Expected result.%s for '%%s' to be '{{.%s}}', but got '%%s'", `+"`{{.Target}}`"+`, result.%s)
		}
`, field, field, field, field, field)
	}

	tmpl += `	}
{{end}}
}
`

	t, err := template.New("testset").Parse(tmpl)
	if err != nil {
		return err
	}

	tests, err := filepath.Glob(filepath.Join("woothee", "testsets", "*.yaml"))
	if err != nil {
		return err
	}

	sort.Strings(tests)
	for _, filename := range tests {
		src, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer src.Close()

		fi, err := src.Stat()
		if err != nil {
			return err
		}

		buf := make([]byte, fi.Size())
		if _, err = io.ReadFull(src, buf); err != nil {
			return err
		}

		var data struct {
			TestName string
			Tests    []*struct {
				Target    string
				Name      string
				Os        string
				OsVersion string `yaml:"os_version,omitempty"`
				Category  string
				Version   string
			}
		}

		if err = yaml.Unmarshal(buf, &data.Tests); err != nil {
			return err
		}

		testname := filepath.Base(filename)
		testname = strings.Replace(testname, "_windows", "_win", -1)
		testname = strings.TrimSuffix(testname, ".yaml")

		data.TestName = testname

		dstFile := fmt.Sprintf("%s_test.go", testname)
		dst, err := os.OpenFile(dstFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer dst.Close()

		for _, d := range data.Tests {
			if d.Name == "" {
				d.Name = "UNKNOWN"
			}
			if d.Category == "" {
				d.Category = "UNKNOWN"
			}
			if d.Os == "" {
				d.Os = "UNKNOWN"
			}
			if d.Version == "" {
				d.Version = "UNKNOWN"
			}
			if d.OsVersion == "" {
				d.OsVersion = "UNKNOWN"
			}
			if d.Version == "" {
				d.Version = "UNKNOWN"
			}
		}

		if err = t.Execute(dst, data); err != nil {
			return err
		}
	}

	return nil
}
