/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_mobilephone_docomo(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`DoCoMo/2.0 SH01A(c100;TB;W24H16)`)
    if err != nil {
        t.Errorf(`Failed to parse 'DoCoMo/2.0 SH01A(c100;TB;W24H16)': %s`, err)
    } else {
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Version != "SH01A" {
            t.Errorf("Expected result.Version to be 'SH01A', but got '%s'", result.Version)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`DoCoMo/2.0 N07B(c500;TB;W24H16)`)
    if err != nil {
        t.Errorf(`Failed to parse 'DoCoMo/2.0 N07B(c500;TB;W24H16)': %s`, err)
    } else {
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Version != "N07B" {
            t.Errorf("Expected result.Version to be 'N07B', but got '%s'", result.Version)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)': %s`, err)
    } else {
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Version != "D505i" {
            t.Errorf("Expected result.Version to be 'D505i', but got '%s'", result.Version)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`DoCoMo/1.0/N505i/c20/TB/W24H12`)
    if err != nil {
        t.Errorf(`Failed to parse 'DoCoMo/1.0/N505i/c20/TB/W24H12': %s`, err)
    } else {
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Version != "N505i" {
            t.Errorf("Expected result.Version to be 'N505i', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)': %s`, err)
    } else {
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
        if result.Version != "N905imyu_W" {
            t.Errorf("Expected result.Version to be 'N905imyu_W', but got '%s'", result.Version)
        }
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (F02B;FOMA;like Gecko)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (F02B;FOMA;like Gecko)': %s`, err)
    } else {
        if result.Os != "docomo" {
            t.Errorf("Expected result.Os to be 'docomo', but got '%s'", result.Os)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Name != "docomo" {
            t.Errorf("Expected result.Name to be 'docomo', but got '%s'", result.Name)
        }
        if result.Version != "F02B" {
            t.Errorf("Expected result.Version to be 'F02B', but got '%s'", result.Version)
        }
    }
}
