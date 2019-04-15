/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_pc_lowpriority(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`Sleipnir/2.9.9`)
	if err != nil {
		t.Errorf(`Failed to parse 'Sleipnir/2.9.9': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Sleipnir/2.9.9`, result.Category)
		}
		if result.Name != "Sleipnir" {
			t.Errorf("Expected result.Name for '%s' to be 'Sleipnir', but got '%s'", `Sleipnir/2.9.9`, result.Name)
		}
		if true && result.Os != "Windows UNKNOWN Ver" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows UNKNOWN Ver', but got '%s'", `Sleipnir/2.9.9`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Sleipnir/2.9.9`, result.OsVersion)
		}
		if true && result.Version != "2.9.9" {
			t.Errorf("Expected result.Version for '%s' to be '2.9.9', but got '%s'", `Sleipnir/2.9.9`, result.Version)
		}
	}

}
