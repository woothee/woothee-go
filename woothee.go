package woothee

import (
  "errors"
)

type Result struct {
  Name      string  `json:"name"`
  Category  string  `json:"category"`
  Os        string  `json:"os"`
  Type      string  `json:"type"`
  Version   string  `json:"version"`
  Vendor    string  `json:"vendor"`
}

type DataSet map[string]*Result

const (
  VALUE_UNKNOWN         = "UNKNOWN"
  CATEGORY_PC           = "pc"
  CATEGORY_SMARTPHONE   = "smartphone"
  CATEGORY_MOBILEPHONE  = "mobilephone"
  CATEGORY_APPLIANCE    = "appliance"
  CATEGORY_CRAWLER      = "crawler"
  CATEGORY_MISC         = "misc"
)

var (
  EmptyResult   = &Result { VALUE_UNKNOWN, VALUE_UNKNOWN, VALUE_UNKNOWN, VALUE_UNKNOWN, VALUE_UNKNOWN, VALUE_UNKNOWN }
  ErrNoMatch    = errors.New("No match")
  ErrNoDataSet  = errors.New("No such dataset")
  DefaultParser = NewParser()
)

func (r *Result) Clone() (*Result) {
  return &Result { r.Name, r.Category, r.Os, r.Type, r.Version, r.Vendor }
}

func Parse(agent string) (*Result, error) {
  return DefaultParser.Parse(agent)
}