/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_smartphone_misc(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0': %s`, err)
    } else {
        if result.Category != "smartphone" {
            t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
        }
        if result.Version != "18.0" {
            t.Errorf("Expected result.Version to be '18.0', but got '%s'", result.Version)
        }
        if result.Name != "Firefox" {
            t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
        }
        if result.Os != "Firefox OS" {
            t.Errorf("Expected result.Os to be 'Firefox OS', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)': %s`, err)
    } else {
        if result.Os != "Windows Phone OS" {
            t.Errorf("Expected result.Os to be 'Windows Phone OS', but got '%s'", result.Os)
        }
        if result.Category != "smartphone" {
            t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
        }
        if result.Version != "7.0" {
            t.Errorf("Expected result.Version to be '7.0', but got '%s'", result.Version)
        }
        if result.Name != "Internet Explorer" {
            t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT': %s`, err)
    } else {
        if result.Category != "smartphone" {
            t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
        }
        if result.Name != "Internet Explorer" {
            t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
        }
        if result.Version != "6.0" {
            t.Errorf("Expected result.Version to be '6.0', but got '%s'", result.Version)
        }
        if result.Os != "Windows CE" {
            t.Errorf("Expected result.Os to be 'Windows CE', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`)
    if err != nil {
        t.Errorf(`Failed to parse 'Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54': %s`, err)
    } else {
        if result.Name != "Opera" {
            t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
        }
        if result.Version != "10.54" {
            t.Errorf("Expected result.Version to be '10.54', but got '%s'", result.Version)
        }
        if result.Category != "smartphone" {
            t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
        }
        if result.Os != "BlackBerry" {
            t.Errorf("Expected result.Os to be 'BlackBerry', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`)
    if err != nil {
        t.Errorf(`Failed to parse 'BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220': %s`, err)
    } else {
        if result.Os != "BlackBerry" {
            t.Errorf("Expected result.Os to be 'BlackBerry', but got '%s'", result.Os)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
        if result.Category != "smartphone" {
            t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
        }
    }
}
