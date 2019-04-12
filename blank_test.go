/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_blank(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(``)
	if err != nil {
		t.Errorf(`Failed to parse '': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Os)
		}
		if true && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", ``, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Version)
		}
	}
	result, err = Parse(``)
	if err != nil {
		t.Errorf(`Failed to parse '': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Os)
		}
		if true && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", ``, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", ``, result.Version)
		}
	}
	result, err = Parse(`-`)
	if err != nil {
		t.Errorf(`Failed to parse '-': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category for '%s' to be 'UNKNOWN', but got '%s'", `-`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `-`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `-`, result.Os)
		}
		if true && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `-`, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `-`, result.Version)
		}
	}

}
