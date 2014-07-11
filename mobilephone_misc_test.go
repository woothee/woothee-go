/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_mobilephone_misc(t *testing.T) {
	var result *Result
	var err error

	result, err = Parse(`Mozilla/5.0 (jig browser core; SH03B)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (jig browser core; SH03B)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name to be 'jig browser', but got '%s'", result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os to be 'jig', but got '%s'", result.Os)
		}
		if result.Version != "SH03B" {
			t.Errorf("Expected result.Version to be 'SH03B', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name to be 'jig browser', but got '%s'", result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os to be 'jig', but got '%s'", result.Os)
		}
		if result.Version != "F10B" {
			t.Errorf("Expected result.Version to be 'F10B', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (jig browser9)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (jig browser9)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name to be 'jig browser', but got '%s'", result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os to be 'jig', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name to be 'emobile', but got '%s'", result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os to be 'emobile', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name to be 'emobile', but got '%s'", result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os to be 'emobile', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name to be 'emobile', but got '%s'", result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os to be 'emobile', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`)
	if err != nil {
		t.Errorf(`Failed to parse 'Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name to be 'emobile', but got '%s'", result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os to be 'emobile', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "SymbianOS" {
			t.Errorf("Expected result.Os to be 'SymbianOS', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "Mobile Transcoder" {
			t.Errorf("Expected result.Name to be 'Mobile Transcoder', but got '%s'", result.Name)
		}
		if result.Os != "Mobile Transcoder" {
			t.Errorf("Expected result.Os to be 'Mobile Transcoder', but got '%s'", result.Os)
		}
		if result.Version != "Google" {
			t.Errorf("Expected result.Version to be 'Google', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category to be 'mobilephone', but got '%s'", result.Category)
		}
		if result.Name != "Mobile Transcoder" {
			t.Errorf("Expected result.Name to be 'Mobile Transcoder', but got '%s'", result.Name)
		}
		if result.Os != "Mobile Transcoder" {
			t.Errorf("Expected result.Os to be 'Mobile Transcoder', but got '%s'", result.Os)
		}
		if result.Version != "livedoor" {
			t.Errorf("Expected result.Version to be 'livedoor', but got '%s'", result.Version)
		}
	}
}
