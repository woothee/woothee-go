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
			t.Errorf("Expected result.Category to be 'UNKNOWN', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(``)
	if err != nil {
		t.Errorf(`Failed to parse '': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category to be 'UNKNOWN', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(``)
	if err != nil {
		t.Errorf(`Failed to parse '': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category to be 'UNKNOWN', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(``)
	if err != nil {
		t.Errorf(`Failed to parse '': %s`, err)
	} else {
		if result.Category != "UNKNOWN" {
			t.Errorf("Expected result.Category to be 'UNKNOWN', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
}
