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
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Os)
		}
		if true && result.OsVersion != "5.0.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '5.0.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.OsVersion)
		}
		if true && result.Version != "5.1" {
			t.Errorf("Expected result.Version for '%s' to be '5.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`, result.Os)
		}
		if true && result.OsVersion != "8.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`, result.OsVersion)
		}
		if true && result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`, result.Category)
		}
		if result.Name != "Webview" {
			t.Errorf("Expected result.Name for '%s' to be 'Webview', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`, result.Os)
		}
		if true && result.OsVersion != "8.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B411`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`, result.Name)
		}
		if true && result.Os != "iPad" {
			t.Errorf("Expected result.Os for '%s' to be 'iPad', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`, result.Os)
		}
		if true && result.OsVersion != "4.3.2" {
			t.Errorf("Expected result.OsVersion for '%s' to be '4.3.2', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`, result.OsVersion)
		}
		if true && result.Version != "5.0.2" {
			t.Errorf("Expected result.Version for '%s' to be '5.0.2', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_2 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`, result.Category)
		}
		if result.Name != "Webview" {
			t.Errorf("Expected result.Name for '%s' to be 'Webview', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`, result.Name)
		}
		if true && result.Os != "iPad" {
			t.Errorf("Expected result.Os for '%s' to be 'iPad', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`, result.Os)
		}
		if true && result.OsVersion != "4.3.5" {
			t.Errorf("Expected result.OsVersion for '%s' to be '4.3.5', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`, result.OsVersion)
		}
		if true && result.Version != "5.0.2" {
			t.Errorf("Expected result.Version for '%s' to be '5.0.2', but got '%s'", `Mozilla/5.0 (iPad; U; CPU OS 4_3_5 like Mac OS X; ja-jp) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8L1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Name)
		}
		if true && result.Os != "iPod" {
			t.Errorf("Expected result.Os for '%s' to be 'iPod', but got '%s'", `Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Os)
		}
		if true && result.OsVersion != "5.0.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '5.0.1', but got '%s'", `Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.OsVersion)
		}
		if true && result.Version != "5.1" {
			t.Errorf("Expected result.Version for '%s' to be '5.1', but got '%s'", `Mozilla/5.0 (iPod; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`, result.Name)
		}
		if true && result.Os != "iPod" {
			t.Errorf("Expected result.Os for '%s' to be 'iPod', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`, result.Os)
		}
		if true && result.OsVersion != "7.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be '7.0', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`, result.OsVersion)
		}
		if true && result.Version != "7.0" {
			t.Errorf("Expected result.Version for '%s' to be '7.0', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`, result.Category)
		}
		if result.Name != "Webview" {
			t.Errorf("Expected result.Name for '%s' to be 'Webview', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`, result.Name)
		}
		if true && result.Os != "iPod" {
			t.Errorf("Expected result.Os for '%s' to be 'iPod', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`, result.Os)
		}
		if true && result.OsVersion != "8.1.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.1.3', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (iPod touch; CPU iPhone OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12B466`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`, result.Os)
		}
		if true && result.OsVersion != "5.1.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '5.1.1', but got '%s'", `Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`, result.OsVersion)
		}
		if true && result.Version != "19.0.1084.60" {
			t.Errorf("Expected result.Version for '%s' to be '19.0.1084.60', but got '%s'", `Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; ja-jp) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`, result.Os)
		}
		if true && result.OsVersion != "8.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.3', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`, result.OsVersion)
		}
		if true && result.Version != "1.0" {
			t.Errorf("Expected result.Version for '%s' to be '1.0', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`, result.Category)
		}
		if result.Name != "Edge" {
			t.Errorf("Expected result.Name for '%s' to be 'Edge', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`, result.Os)
		}
		if true && result.OsVersion != "10.3.2" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.3.2', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`, result.OsVersion)
		}
		if true && result.Version != "41.1.35.1" {
			t.Errorf("Expected result.Version for '%s' to be '41.1.35.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 Safari/603.2.4 EdgiOS/41.1.35.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`, result.Category)
		}
		if result.Name != "Google Search App" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Search App', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`, result.Os)
		}
		if true && result.OsVersion != "11.1.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '11.1.1', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`, result.OsVersion)
		}
		if true && result.Version != "41.0.178428663" {
			t.Errorf("Expected result.Version for '%s' to be '41.0.178428663', but got '%s'", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) GSA/41.0.178428663 Mobile/15B150 Safari/604.1`, result.Version)
		}
	}
	result, err = Parse(`Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`, result.Name)
		}
		if true && result.Os != "iPhone" {
			t.Errorf("Expected result.Os for '%s' to be 'iPhone', but got '%s'", `Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Girls/2.0 (livedoor Co.,Ltd.; Peachy 2.1; iPhone; RSS Version 2.0; +http://girls.livedoor.com/)`, result.Version)
		}
	}
	result, err = Parse(`MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`, result.Name)
		}
		if true && result.Os != "iOS" {
			t.Errorf("Expected result.Os for '%s' to be 'iOS', but got '%s'", `MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`, result.OsVersion)
		}
		if true && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `MobileSafari/7534.48.3 CFNetwork/548.0.4 Darwin/11.0.0`, result.Version)
		}
	}
	result, err = Parse(`Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0': %s`, err)
	} else {
		if result.Category != "smartphone" {
			t.Errorf("Expected result.Category for '%s' to be 'smartphone', but got '%s'", `Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`, result.Name)
		}
		if true && result.Os != "iOS" {
			t.Errorf("Expected result.Os for '%s' to be 'iOS', but got '%s'", `Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Blogos/1.13 CFNetwork/548.0.4 Darwin/11.0.0`, result.Version)
		}
	}

}
