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
  CATEGORY_PC           = "pc"
  CATEGORY_SMARTPHONE   = "smartphone"
  CATEGORY_MOBILEPHONE  = "mobilephone"
  CATEGORY_APPLIANCE    = "appliance"
  CATEGORY_CRAWLER      = "crawler"
  CATEGORY_MISC         = "misc"
  CATEGORY_UNKNOWN      = "UNKNOWN"
)

var (
  EmptyResult   = &Result { Category: CATEGORY_UNKNOWN }
  ErrNoMatch    = errors.New("No match")
  ErrNoDataSet  = errors.New("No such dataset")
  DefaultParser = NewParser()
)

func Parse(agent string) (result *Result, err error) {
  if agent == "" || agent == "-" {
    result = EmptyResult
  }

  result, err = DefaultParser.TryCrawler(agent)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}