package woothee

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestParse(t *testing.T) {
	err := filepath.Walk("./woothee/testsets", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".yaml") {
			src, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file: %s: %s", path, err)
			}
			defer src.Close()

			buf, err := ioutil.ReadAll(src)
			if err != nil {
				return fmt.Errorf("failed to read file: %s: %s", path, err)
			}

			var data []struct {
				Target    string `yaml:"target"`
				Name      string `yaml:"name"`
				Category  string `yaml:"category"`
				Os        string `yaml:"os,omitempty"`
				Version   string `yaml:"version,omitempty"`
				OsVersion string `yaml:"os_version,omitempty"`
			}

			if err := yaml.Unmarshal(buf, &data); err != nil {
				return fmt.Errorf("failed to unmarshal YAML: %s: %s", path, err)
			}

			for _, rec := range data {
				result, err := Parse(rec.Target)
				if err != nil {
					return fmt.Errorf("failed to parse user  agent: %s: %s, %s", path, rec.Target, err)
				}

				if rec.Name != result.Name {
					t.Errorf("name not matched: %s: expected: %s: actual: %s", rec.Target, rec.Name, result.Name)
				}
				if rec.Category != result.Category {
					t.Errorf("category not matched: %s: expected: %s: actual: %s", rec.Target, rec.Category, result.Category)
				}
				if rec.Os != "" && rec.Os != result.Os {
					t.Errorf("os not matched: %s: expected: %s: actual: %s", rec.Target, rec.Os, result.Os)
				}
				if rec.OsVersion != "" && rec.OsVersion != result.OsVersion {
					t.Errorf("os_version not matched: %s: expected: %s: actual: %s", rec.Target, rec.OsVersion, result.OsVersion)
				}
				if rec.Version != "" && rec.Version != result.Version {
					t.Errorf("version not matched: %s: expected: %s: actual: %s", rec.Target, rec.Version, result.Version)
				}
			}
		}

		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk files: %s", err)
	}
}
