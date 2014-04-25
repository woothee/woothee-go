/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_pc_windows(t *testing.T) {
	var result *Result
	var err error

	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)': %s`, err)
	} else {
		if result.Version != "6.0" {
			t.Errorf("Expected result.Version to be '6.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows 2000" {
			t.Errorf("Expected result.Os to be 'Windows 2000', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Version != "7.0" {
			t.Errorf("Expected result.Version to be '7.0', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version to be '8.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)': %s`, err)
	} else {
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version to be '8.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os to be 'Windows Vista', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)': %s`, err)
	} else {
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version to be '8.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os to be 'Windows 7', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)': %s`, err)
	} else {
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Version != "9.0" {
			t.Errorf("Expected result.Version to be '9.0', but got '%s'", result.Version)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os to be 'Windows 7', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)': %s`, err)
	} else {
		if result.Os != "Windows 8" {
			t.Errorf("Expected result.Os to be 'Windows 8', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Version != "10.0" {
			t.Errorf("Expected result.Version to be '10.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version to be '11.0', but got '%s'", result.Version)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os to be 'Windows 8.1', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/4.78 [ja] (Win98; U)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.78 [ja] (Win98; U)': %s`, err)
	} else {
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows 98" {
			t.Errorf("Expected result.Os to be 'Windows 98', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)': %s`, err)
	} else {
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os to be 'Windows 7', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13': %s`, err)
	} else {
		if result.Version != "0.2.149.27" {
			t.Errorf("Expected result.Version to be '0.2.149.27', but got '%s'", result.Version)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name to be 'Chrome', but got '%s'", result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1': %s`, err)
	} else {
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os to be 'Windows Vista', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Version != "9.0.1" {
			t.Errorf("Expected result.Version to be '9.0.1', but got '%s'", result.Version)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7': %s`, err)
	} else {
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
		}
		if result.Version != "5.1.2" {
			t.Errorf("Expected result.Version to be '5.1.2', but got '%s'", result.Version)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Opera/9.52 (Windows NT 5.1; U; ja)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.52 (Windows NT 5.1; U; ja)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Version != "9.52" {
			t.Errorf("Expected result.Version to be '9.52', but got '%s'", result.Version)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64': %s`, err)
	} else {
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os to be 'Windows 8.1', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Version != "20.0.1387.64" {
			t.Errorf("Expected result.Version to be '20.0.1387.64', but got '%s'", result.Version)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0': %s`, err)
	} else {
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
		}
		if result.Version != "6.0" {
			t.Errorf("Expected result.Version to be '6.0', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)': %s`, err)
	} else {
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os to be 'Windows Vista', but got '%s'", result.Os)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0': %s`, err)
	} else {
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
		}
		if result.Category != "pc" {
			t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os to be 'Windows XP', but got '%s'", result.Os)
		}
	}
}
