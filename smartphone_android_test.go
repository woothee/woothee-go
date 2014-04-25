/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_smartphone_android(t *testing.T) {
	var result *Result
	var err error

	result, err = Parse(`Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Version != "4.0" {
			t.Errorf("Expected result.Version to be '4.0', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13': %s`, err)
	} else {
		if result.Version != "4.0" {
			t.Errorf("Expected result.Version to be '4.0', but got '%s'", result.Version)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7': %s`, err)
	} else {
		if result.Version != "16.0.912.75" {
			t.Errorf("Expected result.Version to be '16.0.912.75', but got '%s'", result.Version)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name to be 'Chrome', but got '%s'", result.Name)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
		if result.Version != "10.54" {
			t.Errorf("Expected result.Version to be '10.54', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172': %s`, err)
	} else {
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
		if result.Version != "20.0.1396.73172" {
			t.Errorf("Expected result.Version to be '20.0.1396.73172', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0': %s`, err)
	} else {
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Version != "14.0" {
			t.Errorf("Expected result.Version to be '14.0', but got '%s'", result.Version)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0': %s`, err)
	} else {
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
		}
		if result.Version != "14.0" {
			t.Errorf("Expected result.Version to be '14.0', but got '%s'", result.Version)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)': %s`, err)
	} else {
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`LDNReader/2.0.1 (Android)`)
	if err != nil {
		t.Errorf(`Failed to parse 'LDNReader/2.0.1 (Android)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category to be 'smartphone', but got '%s'", result.Category)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os to be 'Android', but got '%s'", result.Os)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
	}
}
