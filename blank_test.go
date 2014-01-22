/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_blank(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(``)
    if err != nil {
        t.Errorf(`Failed to parse '': %s`, err)
    } else {
        if result.Name != "" {
            t.Errorf("Expected result.Name to be '', but got '%s'", result.Name)
        }
        if result.Version != "" {
            t.Errorf("Expected result.Version to be '', but got '%s'", result.Version)
        }
        if result.Category != "" {
            t.Errorf("Expected result.Category to be '', but got '%s'", result.Category)
        }
        if result.Os != "" {
            t.Errorf("Expected result.Os to be '', but got '%s'", result.Os)
        }
    }
    result, err = Parse(``)
    if err != nil {
        t.Errorf(`Failed to parse '': %s`, err)
    } else {
        if result.Category != "" {
            t.Errorf("Expected result.Category to be '', but got '%s'", result.Category)
        }
        if result.Name != "" {
            t.Errorf("Expected result.Name to be '', but got '%s'", result.Name)
        }
        if result.Version != "" {
            t.Errorf("Expected result.Version to be '', but got '%s'", result.Version)
        }
        if result.Os != "" {
            t.Errorf("Expected result.Os to be '', but got '%s'", result.Os)
        }
    }
    result, err = Parse(``)
    if err != nil {
        t.Errorf(`Failed to parse '': %s`, err)
    } else {
        if result.Os != "" {
            t.Errorf("Expected result.Os to be '', but got '%s'", result.Os)
        }
        if result.Name != "" {
            t.Errorf("Expected result.Name to be '', but got '%s'", result.Name)
        }
        if result.Version != "" {
            t.Errorf("Expected result.Version to be '', but got '%s'", result.Version)
        }
        if result.Category != "" {
            t.Errorf("Expected result.Category to be '', but got '%s'", result.Category)
        }
    }
    result, err = Parse(``)
    if err != nil {
        t.Errorf(`Failed to parse '': %s`, err)
    } else {
        if result.Os != "" {
            t.Errorf("Expected result.Os to be '', but got '%s'", result.Os)
        }
        if result.Name != "" {
            t.Errorf("Expected result.Name to be '', but got '%s'", result.Name)
        }
        if result.Version != "" {
            t.Errorf("Expected result.Version to be '', but got '%s'", result.Version)
        }
        if result.Category != "" {
            t.Errorf("Expected result.Category to be '', but got '%s'", result.Category)
        }
    }
}
