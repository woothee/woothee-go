/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_mobilephone_willcom(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`Mozilla/3.0(WILLCOM;TOSHIBA/WX320T/2;1/1/C128) NetFront/3.4`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0(WILLCOM;TOSHIBA/WX320T/2;1/1/C128) NetFront/3.4': %s`, err)
    } else {
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Os != "WILLCOM" {
            t.Errorf("Expected result.Os to be 'WILLCOM', but got '%s'", result.Os)
        }
        if result.Version != "WX320T" {
            t.Errorf("Expected result.Version to be 'WX320T', but got '%s'", result.Version)
        }
        if result.Name != "WILLCOM" {
            t.Errorf("Expected result.Name to be 'WILLCOM', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/3.0(WILLCOM;SANYO/WX310SA/2;1/1/C128) NetFront/3.3`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0(WILLCOM;SANYO/WX310SA/2;1/1/C128) NetFront/3.3': %s`, err)
    } else {
        if result.Version != "WX310SA" {
            t.Errorf("Expected result.Version to be 'WX310SA', but got '%s'", result.Version)
        }
        if result.Name != "WILLCOM" {
            t.Errorf("Expected result.Name to be 'WILLCOM', but got '%s'", result.Name)
        }
        if result.Os != "WILLCOM" {
            t.Errorf("Expected result.Os to be 'WILLCOM', but got '%s'", result.Os)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/3.0(WILLCOM;KYOCERA/WX331K/2;1.0.8.13.000000/0.1/C100) Opera 7.2 EX`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0(WILLCOM;KYOCERA/WX331K/2;1.0.8.13.000000/0.1/C100) Opera 7.2 EX': %s`, err)
    } else {
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Os != "WILLCOM" {
            t.Errorf("Expected result.Os to be 'WILLCOM', but got '%s'", result.Os)
        }
        if result.Version != "WX331K" {
            t.Errorf("Expected result.Version to be 'WX331K', but got '%s'", result.Version)
        }
        if result.Name != "Opera" {
            t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/3.0(DDIPOCKET;JRC/AH-J3001V,AH-J3002V/1.0/0100/c50)CNF/2.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0(DDIPOCKET;JRC/AH-J3001V,AH-J3002V/1.0/0100/c50)CNF/2.0': %s`, err)
    } else {
        if result.Name != "WILLCOM" {
            t.Errorf("Expected result.Name to be 'WILLCOM', but got '%s'", result.Name)
        }
        if result.Version != "AH-J3001V,AH-J3002V" {
            t.Errorf("Expected result.Version to be 'AH-J3001V,AH-J3002V', but got '%s'", result.Version)
        }
        if result.Os != "WILLCOM" {
            t.Errorf("Expected result.Os to be 'WILLCOM', but got '%s'", result.Os)
        }
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/3.0(DDIPOCKET;KYOCERA/AH-K3001V/1.4.1.67.000000/0.1/C100) Opera 7.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0(DDIPOCKET;KYOCERA/AH-K3001V/1.4.1.67.000000/0.1/C100) Opera 7.0': %s`, err)
    } else {
        if result.Category != "mobilephone" {
            t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
        }
        if result.Os != "WILLCOM" {
            t.Errorf("Expected result.Os to be 'WILLCOM', but got '%s'", result.Os)
        }
        if result.Version != "AH-K3001V" {
            t.Errorf("Expected result.Version to be 'AH-K3001V', but got '%s'", result.Version)
        }
        if result.Name != "Opera" {
            t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
        }
    }
}
