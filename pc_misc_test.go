/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_pc_misc(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`, result.Os)
		}
		if result.OsVersion != "10.5.4" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.5.4', but got '%s'", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`, result.OsVersion)
		}
		if result.Version != "3.1.2" {
			t.Errorf("Expected result.Version for '%s' to be '3.1.2', but got '%s'", `Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`, result.Os)
		}
		if result.OsVersion != "10.7.5" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.7.5', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`, result.OsVersion)
		}
		if result.Version != "27.0.1453.93" {
			t.Errorf("Expected result.Version for '%s' to be '27.0.1453.93', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`, result.Os)
		}
		if result.OsVersion != "10.7" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.7', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`, result.OsVersion)
		}
		if result.Version != "21.0" {
			t.Errorf("Expected result.Version for '%s' to be '21.0', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`, result.Version)
		}
	}
	result, err = Parse(`Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`, result.Os)
		}
		if result.OsVersion != "10.8.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.8.3', but got '%s'", `Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`, result.OsVersion)
		}
		if result.Version != "12.15" {
			t.Errorf("Expected result.Version for '%s' to be '12.15', but got '%s'", `Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`, result.Os)
		}
		if result.OsVersion != "10.9.2" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.9.2', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`, result.OsVersion)
		}
		if result.Version != "20.0.1387.82" {
			t.Errorf("Expected result.Version for '%s' to be '20.0.1387.82', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.154 Safari/537.36 OPR/20.0.1387.82`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`, result.Category)
		}
		if result.Name != "Vivaldi" {
			t.Errorf("Expected result.Name for '%s' to be 'Vivaldi', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`, result.Os)
		}
		if result.OsVersion != "10.11.4" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.11.4', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`, result.OsVersion)
		}
		if result.Version != "1.1.453.52" {
			t.Errorf("Expected result.Version for '%s' to be '1.1.453.52', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36 Vivaldi/1.1.453.52`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`, result.Os)
		}
		/* Skip for now
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`, result.OsVersion)
		}*/
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`, result.Name)
		}
		if result.Os != "Mac OSX" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OSX', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`, result.Os)
		}
		if result.OsVersion != "10.5" {
			t.Errorf("Expected result.OsVersion for '%s' to be '10.5', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`, result.Name)
		}
		if result.Os != "Mac OS Classic" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OS Classic', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`, result.Os)
		}
		if result.OsVersion != "9.2.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be '9.2.1', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`, result.OsVersion)
		}
		if result.Version != "3.6.8" {
			t.Errorf("Expected result.Version for '%s' to be '3.6.8', but got '%s'", `Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`, result.Name)
		}
		if result.Os != "Mac OS Classic" {
			t.Errorf("Expected result.Os for '%s' to be 'Mac OS Classic', but got '%s'", `Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`, result.OsVersion)
		}
		if result.Version != "5.17" {
			t.Errorf("Expected result.Version for '%s' to be '5.17', but got '%s'", `Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.OsVersion)
		}
		if result.Version != "9.0.1" {
			t.Errorf("Expected result.Version for '%s' to be '9.0.1', but got '%s'", `Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`, result.OsVersion)
		}
		if result.Version != "16.0.912.75" {
			t.Errorf("Expected result.Version for '%s' to be '16.0.912.75', but got '%s'", `Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`, result.Category)
		}
		if result.Name != "Vivaldi" {
			t.Errorf("Expected result.Name for '%s' to be 'Vivaldi', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`, result.OsVersion)
		}
		if result.Version != "1.0.321.3" {
			t.Errorf("Expected result.Version for '%s' to be '1.0.321.3', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.99 Safari/537.36 Vivaldi/1.0.321.3`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`, result.Name)
		}
		if result.Os != "BSD" {
			t.Errorf("Expected result.Os for '%s' to be 'BSD', but got '%s'", `Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`, result.Os)
		}
		/* Skip for now
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`, result.OsVersion)
		}*/
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`, result.Version)
		}
	}
	result, err = Parse(`Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`, result.Name)
		}
		if result.Os != "BSD" {
			t.Errorf("Expected result.Os for '%s' to be 'BSD', but got '%s'", `Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`, result.Os)
		}
		if result.OsVersion != "8.2-RELEASE-p3 amd64" {
			t.Errorf("Expected result.OsVersion for '%s' to be '8.2-RELEASE-p3 amd64', but got '%s'", `Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`, result.OsVersion)
		}
		if result.Version != "11.52" {
			t.Errorf("Expected result.Version for '%s' to be '11.52', but got '%s'", `Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`, result.Name)
		}
		if result.Os != "ChromeOS" {
			t.Errorf("Expected result.Os for '%s' to be 'ChromeOS', but got '%s'", `Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`, result.Os)
		}
		if result.OsVersion != "x86_64 5116.115.4" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'x86_64 5116.115.4', but got '%s'", `Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`, result.OsVersion)
		}
		if result.Version != "33.0.1750.152" {
			t.Errorf("Expected result.Version for '%s' to be '33.0.1750.152', but got '%s'", `Mozilla/5.0 (X11; CrOS x86_64 5116.115.4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`, result.Category)
		}
		if result.Name != "Yandex Browser" {
			t.Errorf("Expected result.Name for '%s' to be 'Yandex Browser', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`, result.Name)
		}
		if result.Os != "Linux" {
			t.Errorf("Expected result.Os for '%s' to be 'Linux', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`, result.OsVersion)
		}
		if result.Version != "17.11.1.1087" {
			t.Errorf("Expected result.Version for '%s' to be '17.11.1.1087', but got '%s'", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 YaBrowser/17.11.1.1087 (beta) Yowser/2.5 Safari/537.36`, result.Version)
		}
	}
}
