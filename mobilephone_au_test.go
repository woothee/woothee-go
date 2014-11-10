/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_mobilephone_au(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`KDDI-TS3V UP.Browser/6.2_7.2.7.1.K.6.210 (GUI) MMP/2.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'KDDI-TS3V UP.Browser/6.2_7.2.7.1.K.6.210 (GUI) MMP/2.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "au by KDDI" {
			t.Errorf("Expected result.Name to be 'au by KDDI', but got '%s'", result.Name)
		}
		if result.Os != "au" {
			t.Errorf("Expected result.Os to be 'au', but got '%s'", result.Os)
		}
		if result.Version != "TS3V" {
			t.Errorf("Expected result.Version to be 'TS3V', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`KDDI-CA3H UP.Browser/6.2_7.2.7.1.K.5.207 (GUI) MMP/2.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'KDDI-CA3H UP.Browser/6.2_7.2.7.1.K.5.207 (GUI) MMP/2.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "au by KDDI" {
			t.Errorf("Expected result.Name to be 'au by KDDI', but got '%s'", result.Name)
		}
		if result.Os != "au" {
			t.Errorf("Expected result.Os to be 'au', but got '%s'", result.Os)
		}
		if result.Version != "CA3H" {
			t.Errorf("Expected result.Version to be 'CA3H', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 Opera/9.5 (KDDI-SH3F; BREW; Opera Mobi; U; ja) Presto/2.2.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 Opera/9.5 (KDDI-SH3F; BREW; Opera Mobi; U; ja) Presto/2.2.1': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
		if result.Os != "au" {
			t.Errorf("Expected result.Os to be 'au', but got '%s'", result.Os)
		}
		if result.Version != "SH3F" {
			t.Errorf("Expected result.Version to be 'SH3F', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; KDDI-CA3B) Opera 8.60 [ja]`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; KDDI-CA3B) Opera 8.60 [ja]': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Os != "au" {
			t.Errorf("Expected result.Os to be 'au', but got '%s'", result.Os)
		}
		if result.Version != "CA3B" {
			t.Errorf("Expected result.Version to be 'CA3B', but got '%s'", result.Version)
		}
	}

}
