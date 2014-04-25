/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_pc_lowpriority(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`Sleipnir/2.9.9`)
    if err != nil {
        t.Errorf(`Failed to parse 'Sleipnir/2.9.9': %s`, err)
    } else {
        if result.Os != "Windows UNKNOWN Ver" {
            t.Errorf("Expected result.Os to be 'Windows UNKNOWN Ver', but got '%s'", result.Os)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "Sleipnir" {
            t.Errorf("Expected result.Name to be 'Sleipnir', but got '%s'", result.Name)
        }
        if result.Version != "2.9.9" {
            t.Errorf("Expected result.Version to be '2.9.9', but got '%s'", result.Version)
        }
    }
}
