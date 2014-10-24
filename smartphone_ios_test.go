/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_smartphone_ios(t *testing.T) {
	var result *Result
	var err error

	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "iPhone" {
			t.Errorf("Expected result.Os to be 'iPhone', but got '%s'", result.Os)
		}
		if result.OsVersion != "5.0.1" {
			t.Errorf("Expected result.OsVersion to be '5.0.1', but got '%s'", result.OsVersion)
		}
		if result.Version != "5.1" {
			t.Errorf("Expected result.Version to be '5.1', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "iPad" {
			t.Errorf("Expected result.Os to be 'iPad', but got '%s'", result.Os)
		}
		if result.OsVersion != "4.3.2" {
			t.Errorf("Expected result.OsVersion to be '4.3.2', but got '%s'", result.OsVersion)
		}
		if result.Version != "5.0.2" {
			t.Errorf("Expected result.Version to be '5.0.2', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "iPod" {
			t.Errorf("Expected result.Os to be 'iPod', but got '%s'", result.Os)
		}
		if result.OsVersion != "5.0.1" {
			t.Errorf("Expected result.OsVersion to be '5.0.1', but got '%s'", result.OsVersion)
		}
		if result.Version != "5.1" {
			t.Errorf("Expected result.Version to be '5.1', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "iPod" {
			t.Errorf("Expected result.Os to be 'iPod', but got '%s'", result.Os)
		}
		if result.OsVersion != "7.0" {
			t.Errorf("Expected result.OsVersion to be '7.0', but got '%s'", result.OsVersion)
		}
		if result.Version != "7.0" {
			t.Errorf("Expected result.Version to be '7.0', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name to be 'Chrome', but got '%s'", result.Name)
		}
		if result.Os != "iPhone" {
			t.Errorf("Expected result.Os to be 'iPhone', but got '%s'", result.Os)
		}
		if result.OsVersion != "5.1.1" {
			t.Errorf("Expected result.OsVersion to be '5.1.1', but got '%s'", result.OsVersion)
		}
		if result.Version != "19.0.1084.60" {
			t.Errorf("Expected result.Version to be '19.0.1084.60', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Mobile/9A405`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Mobile/9A405': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "iPhone" {
			t.Errorf("Expected result.Os to be 'iPhone', but got '%s'", result.Os)
		}
		if result.OsVersion != "5.0.1" {
			t.Errorf("Expected result.OsVersion to be '5.0.1', but got '%s'", result.OsVersion)
		}
	}
	result, err = Parse(`Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "iPhone" {
			t.Errorf("Expected result.Os to be 'iPhone', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "iOS" {
			t.Errorf("Expected result.Os to be 'iOS', but got '%s'", result.Os)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version to be 'UNKNOWN', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "iOS" {
			t.Errorf("Expected result.Os to be 'iOS', but got '%s'", result.Os)
		}
	}
}
