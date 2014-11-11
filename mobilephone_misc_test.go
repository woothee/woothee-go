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
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (jig browser core; SH03B)`, result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name for '%s' to be 'jig browser', but got '%s'", `Mozilla/5.0 (jig browser core; SH03B)`, result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os for '%s' to be 'jig', but got '%s'", `Mozilla/5.0 (jig browser core; SH03B)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (jig browser core; SH03B)`, result.OsVersion)
		}
		if result.Version != "SH03B" {
			t.Errorf("Expected result.Version for '%s' to be 'SH03B', but got '%s'", `Mozilla/5.0 (jig browser core; SH03B)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`, result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name for '%s' to be 'jig browser', but got '%s'", `Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`, result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os for '%s' to be 'jig', but got '%s'", `Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`, result.OsVersion)
		}
		if result.Version != "F10B" {
			t.Errorf("Expected result.Version for '%s' to be 'F10B', but got '%s'", `Mozilla/4.0 (jig browser9i 1.5.0; F10B; 2004)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (jig browser9)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (jig browser9)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/4.0 (jig browser9)`, result.Category)
		}
		if result.Name != "jig browser" {
			t.Errorf("Expected result.Name for '%s' to be 'jig browser', but got '%s'", `Mozilla/4.0 (jig browser9)`, result.Name)
		}
		if result.Os != "jig" {
			t.Errorf("Expected result.Os for '%s' to be 'jig', but got '%s'", `Mozilla/4.0 (jig browser9)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.0 (jig browser9)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.0 (jig browser9)`, result.Version)
		}
	}
	result, err = Parse(`emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`, result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name for '%s' to be 'emobile', but got '%s'", `emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`, result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os for '%s' to be 'emobile', but got '%s'", `emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `emobile/1.0.0 (H11T; like Gecko; Wireless) NetFront/3.4`, result.Version)
		}
	}
	result, err = Parse(`IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`, result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name for '%s' to be 'emobile', but got '%s'", `IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`, result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os for '%s' to be 'emobile', but got '%s'", `IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `IAC/1.0 (H31IA;like Gecko;OpenBrowser) WWW Browser/ver1.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`, result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name for '%s' to be 'emobile', but got '%s'", `Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`, result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os for '%s' to be 'emobile', but got '%s'", `Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (H11T; like Gecko; OpenBrowser) NetFront/3.4`, result.Version)
		}
	}
	result, err = Parse(`Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`)
	if err != nil {
		t.Errorf(`Failed to parse 'Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`, result.Category)
		}
		if result.Name != "emobile" {
			t.Errorf("Expected result.Name for '%s' to be 'emobile', but got '%s'", `Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`, result.Name)
		}
		if result.Os != "emobile" {
			t.Errorf("Expected result.Os for '%s' to be 'emobile', but got '%s'", `Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Huawei/1.0/H12HW/B000 Browser/Obigo-Browser/Q04A`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`, result.Name)
		}
		if result.Os != "SymbianOS" {
			t.Errorf("Expected result.Os for '%s' to be 'SymbianOS', but got '%s'", `Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`, result.Category)
		}
		if result.Name != "Mobile Transcoder" {
			t.Errorf("Expected result.Name for '%s' to be 'Mobile Transcoder', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`, result.Name)
		}
		if result.Os != "Mobile Transcoder" {
			t.Errorf("Expected result.Os for '%s' to be 'Mobile Transcoder', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`, result.OsVersion)
		}
		if result.Version != "Google" {
			t.Errorf("Expected result.Version for '%s' to be 'Google', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`, result.Category)
		}
		if result.Name != "Mobile Transcoder" {
			t.Errorf("Expected result.Name for '%s' to be 'Mobile Transcoder', but got '%s'", `Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`, result.Name)
		}
		if result.Os != "Mobile Transcoder" {
			t.Errorf("Expected result.Os for '%s' to be 'Mobile Transcoder', but got '%s'", `Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`, result.OsVersion)
		}
		if result.Version != "livedoor" {
			t.Errorf("Expected result.Version for '%s' to be 'livedoor', but got '%s'", `Mozilla/5.0 (compatible; livedoor-Mobile-Gateway/0.02; +http://p.m.livedoor.com/help.html)`, result.Version)
		}
	}

}
