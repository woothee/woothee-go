/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_smartphone_misc(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`, result.Name)
		}
		if result.Os != "Firefox OS" {
			t.Errorf("Expected result.Os for '%s' to be 'Firefox OS', but got '%s'", `Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`, result.Os)
		}
		/* Skip for now
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`, result.OsVersion)
		}*/
		if result.Version != "18.0" {
			t.Errorf("Expected result.Version for '%s' to be '18.0', but got '%s'", `Mozilla/5.0 (Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`, result.Name)
		}
		if result.Os != "Firefox OS" {
			t.Errorf("Expected result.Os for '%s' to be 'Firefox OS', but got '%s'", `Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`, result.Os)
		}
		if result.OsVersion != "26.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be '26.0', but got '%s'", `Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`, result.OsVersion)
		}
		if result.Version != "26.0" {
			t.Errorf("Expected result.Version for '%s' to be '26.0', but got '%s'", `Mozilla/5.0 (Tablet; rv:26.0) Gecko/18.0 Firefox/26.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`, result.Name)
		}
		if result.Os != "Windows Phone OS" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows Phone OS', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`, result.Os)
		}
		if result.OsVersion != "7.5" {
			t.Errorf("Expected result.OsVersion for '%s' to be '7.5', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`, result.OsVersion)
		}
		if result.Version != "7.0" {
			t.Errorf("Expected result.Version for '%s' to be '7.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.5; Trident/3.1; IEMobile/7.0; FujitsuToshibaMobileCommun; IS12T; KDDI)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`, result.Name)
		}
		if result.Os != "Windows Phone OS" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows Phone OS', but got '%s'", `Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`, result.Os)
		}
		if result.OsVersion != "8.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.1', but got '%s'", `Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 930) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`, result.Name)
		}
		if result.Os != "Windows CE" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows CE', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`, result.Os)
		}
		/* Skip for now
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`, result.OsVersion)
		}*/
		if result.Version != "6.0" {
			t.Errorf("Expected result.Version for '%s' to be '6.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.7) S12HT`, result.Version)
		}
	}
	result, err = Parse(`Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Name)
		}
		if result.Os != "BlackBerry" {
			t.Errorf("Expected result.Os for '%s' to be 'BlackBerry', but got '%s'", `Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.OsVersion)
		}
		if result.Version != "10.54" {
			t.Errorf("Expected result.Version for '%s' to be '10.54', but got '%s'", `Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Version)
		}
	}
	result, err = Parse(`BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`)
	if err != nil {
		t.Errorf(`Failed to parse 'BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`, result.Name)
		}
		if result.Os != "BlackBerry" {
			t.Errorf("Expected result.Os for '%s' to be 'BlackBerry', but got '%s'", `BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`, result.Os)
		}
		if result.OsVersion != "5.0.0.1014" {
			t.Errorf("Expected result.OsVersion for '%s' to be '5.0.0.1014', but got '%s'", `BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `BlackBerry9700/5.0.0.1014 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/220`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`, result.Name)
		}
		if result.Os != "BlackBerry 10" {
			t.Errorf("Expected result.Os for '%s' to be 'BlackBerry 10', but got '%s'", `Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`, result.Os)
		}
		if result.OsVersion != "10.3.1.2243" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.3.1.2243', but got '%s'", `Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (BB10; Touch) AppleWebKit/537.35+ (KHTML, like Gecko) Version/10.3.1.2243 Mobile Safari/537.35+`, result.Version)
		}
	}

}
