//go:generate go get gopkg.in/yaml.v2
//go:generate go run maint/gendataset.go
package woothee

import (
	"errors"
)

type Result struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Os        string `json:"os"`
	OsVersion string `json:"os_version"`
	Type      string `json:"type"`
	Version   string `json:"version"`
	Vendor    string `json:"vendor"`
}

type DataSet map[string]*Result

const (
	ValueUnknown        = "UNKNOWN"
	CategoryPC          = "pc"
	CategorySmartphone  = "smartphone"
	CategoryMobilephone = "mobilephone"
	CategoryAppliance   = "appliance"
	CategoryCrawler     = "crawler"
	CategoryMisc        = "misc"
)

var (
	EmptyResult   = &Result{ValueUnknown, ValueUnknown, ValueUnknown, ValueUnknown, ValueUnknown, ValueUnknown, ValueUnknown}
	ErrNoMatch    = errors.New("no match")
	ErrNoDataSet  = errors.New("no such dataset")
	DefaultParser = NewParser()
)

func (r *Result) Clone() *Result {
	return &Result{r.Name, r.Category, r.Os, r.OsVersion, r.Type, r.Version, r.Vendor}
}

// Parse parses the given agent string, and returns a Result struct or
// an error if any
func Parse(agent string) (*Result, error) {
	return DefaultParser.Parse(agent)
}
