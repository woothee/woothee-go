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
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`, result.Os)
		}
		if result.OsVersion != "2.3.5" {
			t.Errorf("Expected result.OsVersion for '%s' to be '2.3.5', but got '%s'", `Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`, result.OsVersion)
		}
		if result.Version != "4.0" {
			t.Errorf("Expected result.Version for '%s' to be '4.0', but got '%s'", `Mozilla/5.0 (Linux; U; Android 2.3.5; ja-jp; ISW11F Build/FGK500) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`, result.Os)
		}
		if result.OsVersion != "3.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '3.1', but got '%s'", `Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`, result.OsVersion)
		}
		if result.Version != "4.0" {
			t.Errorf("Expected result.Version for '%s' to be '4.0', but got '%s'", `Mozilla/5.0 (Linux; U; Android 3.1; ja-jp; L-06C Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`, result.Os)
		}
		if result.OsVersion != "4.0.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be '4.0.3', but got '%s'", `Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`, result.OsVersion)
		}
		if result.Version != "16.0.912.75" {
			t.Errorf("Expected result.Version for '%s' to be '16.0.912.75', but got '%s'", `Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`, result.Version)
		}
	}
	result, err = Parse(`Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.OsVersion)
		}
		if result.Version != "10.54" {
			t.Errorf("Expected result.Version for '%s' to be '10.54', but got '%s'", `Opera/9.80 (Android; Opera Mini/6.5.27452/26.1305; U; ja) Presto/2.8.119 Version/10.54`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`, result.Os)
		}
		if result.OsVersion != "4.2.2" {
			t.Errorf("Expected result.OsVersion for '%s' to be '4.2.2', but got '%s'", `Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`, result.OsVersion)
		}
		if result.Version != "20.0.1396.73172" {
			t.Errorf("Expected result.Version for '%s' to be '20.0.1396.73172', but got '%s'", `Mozilla/5.0 (Linux; Android 4.2.2; SO-01F Build/14.1.H.1.281) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Mobile Safari/537.36 OPR/20.0.1396.73172`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`, result.OsVersion)
		}
		if result.Version != "14.0" {
			t.Errorf("Expected result.Version for '%s' to be '14.0', but got '%s'", `Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`, result.OsVersion)
		}
		if result.Version != "14.0" {
			t.Errorf("Expected result.Version for '%s' to be '14.0', but got '%s'", `Mozilla/5.0 (Android; Tablet; rv:14.0) Gecko/14.0 Firefox/14.0`, result.Version)
		}
	}
	result, err = Parse(`Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`, result.Os)
		}
		if result.OsVersion != "2.3.4" {
			t.Errorf("Expected result.OsVersion for '%s' to be '2.3.4', but got '%s'", `Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Dalvik/1.4.0 (Linux; U; Android 2.3.4; SBM009SH Build/S0008)`, result.Version)
		}
	}
	result, err = Parse(`LDNReader/2.0.1 (Android)`)
	if err != nil {
		t.Errorf(`Failed to parse 'LDNReader/2.0.1 (Android)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `LDNReader/2.0.1 (Android)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `LDNReader/2.0.1 (Android)`, result.Name)
		}
		if result.Os != "Android" {
			t.Errorf("Expected result.Os for '%s' to be 'Android', but got '%s'", `LDNReader/2.0.1 (Android)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `LDNReader/2.0.1 (Android)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `LDNReader/2.0.1 (Android)`, result.Version)
		}
	}

}
