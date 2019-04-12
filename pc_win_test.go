/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_pc_win(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`, result.Name)
		}
		if result.Os != "Windows 2000" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 2000', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`, result.Os)
		}
		if result.OsVersion != "NT 5.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`, result.OsVersion)
		}
		if result.Version != "6.0" {
			t.Errorf("Expected result.Version for '%s' to be '6.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; InfoPath.1)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`, result.OsVersion)
		}
		if result.Version != "7.0" {
			t.Errorf("Expected result.Version for '%s' to be '7.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`, result.OsVersion)
		}
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; 2004/11/08; GoogleT5; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`, result.OsVersion)
		}
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; BOIE8;ENUSMSCOM)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`, result.Name)
		}
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows Vista', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`, result.Os)
		}
		if result.OsVersion != "NT 6.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`, result.OsVersion)
		}
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`, result.OsVersion)
		}
		if result.Version != "8.0" {
			t.Errorf("Expected result.Version for '%s' to be '8.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`, result.OsVersion)
		}
		if result.Version != "9.0" {
			t.Errorf("Expected result.Version for '%s' to be '9.0', but got '%s'", `Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`, result.Name)
		}
		if result.Os != "Windows 8" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 8', but got '%s'", `Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`, result.Os)
		}
		if result.OsVersion != "NT 6.2" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.2', but got '%s'", `Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`, result.OsVersion)
		}
		if result.Version != "10.0" {
			t.Errorf("Expected result.Version for '%s' to be '10.0', but got '%s'", `Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; BOIE9;JAJP; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 8.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.3', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Trident/7.0; MALCJS; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; MASPJS; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 8.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.3', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Name)
		}
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 8.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Os)
		}
		if result.OsVersion != "NT 6.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.3', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.OsVersion)
		}
		if result.Version != "11.0" {
			t.Errorf("Expected result.Version for '%s' to be '11.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; Touch; rv:11.0) like Gecko`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.78 [ja] (Win98; U)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.78 [ja] (Win98; U)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.78 [ja] (Win98; U)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.78 [ja] (Win98; U)`, result.Name)
		}
		if result.Os != "Windows 98" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 98', but got '%s'", `Mozilla/4.78 [ja] (Win98; U)`, result.Os)
		}
		if result.OsVersion != "98" {
			t.Errorf("Expected result.OsVersion for '%s' to be '98', but got '%s'", `Mozilla/4.78 [ja] (Win98; U)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.78 [ja] (Win98; U)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`, result.Category)
		}
		if result.Name != "Edge" {
			t.Errorf("Expected result.Name for '%s' to be 'Edge', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`, result.Name)
		}
		if result.Os != "Windows 10" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 10', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`, result.Os)
		}
		if result.OsVersion != "NT 10.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 10.0', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`, result.Category)
		}
		if result.Name != "Edge" {
			t.Errorf("Expected result.Name for '%s' to be 'Edge', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`, result.Name)
		}
		if result.Os != "Windows 10" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 10', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`, result.Os)
		}
		if result.OsVersion != "NT 10.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 10.0', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`, result.Category)
		}
		if result.Name != "Chrome" {
			t.Errorf("Expected result.Name for '%s' to be 'Chrome', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`, result.OsVersion)
		}
		if result.Version != "0.2.149.27" {
			t.Errorf("Expected result.Version for '%s' to be '0.2.149.27', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Category)
		}
		if result.Name != "Firefox" {
			t.Errorf("Expected result.Name for '%s' to be 'Firefox', but got '%s'", `Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Name)
		}
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows Vista', but got '%s'", `Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Os)
		}
		if result.OsVersion != "NT 6.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.0', but got '%s'", `Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.OsVersion)
		}
		if result.Version != "9.0.1" {
			t.Errorf("Expected result.Version for '%s' to be '9.0.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.0; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`, result.Category)
		}
		if result.Name != "Safari" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari', but got '%s'", `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`, result.OsVersion)
		}
		if result.Version != "5.1.2" {
			t.Errorf("Expected result.Version for '%s' to be '5.1.2', but got '%s'", `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7`, result.Version)
		}
	}
	result, err = Parse(`Opera/9.52 (Windows NT 5.1; U; ja)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.52 (Windows NT 5.1; U; ja)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Opera/9.52 (Windows NT 5.1; U; ja)`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Opera/9.52 (Windows NT 5.1; U; ja)`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Opera/9.52 (Windows NT 5.1; U; ja)`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Opera/9.52 (Windows NT 5.1; U; ja)`, result.OsVersion)
		}
		if result.Version != "9.52" {
			t.Errorf("Expected result.Version for '%s' to be '9.52', but got '%s'", `Opera/9.52 (Windows NT 5.1; U; ja)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`, result.Category)
		}
		if result.Name != "Vivaldi" {
			t.Errorf("Expected result.Name for '%s' to be 'Vivaldi', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`, result.Name)
		}
		if result.Os != "Windows 10" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 10', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`, result.Os)
		}
		if result.OsVersion != "NT 10.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 10.0', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`, result.OsVersion)
		}
		if result.Version != "1.0.380.2" {
			t.Errorf("Expected result.Version for '%s' to be '1.0.380.2', but got '%s'", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.88 Safari/537.36 Vivaldi/1.0.380.2`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`, result.Category)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name for '%s' to be 'Opera', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`, result.Name)
		}
		if result.Os != "Windows 8.1" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 8.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`, result.Os)
		}
		if result.OsVersion != "NT 6.3" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.3', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`, result.OsVersion)
		}
		if result.Version != "20.0.1387.64" {
			t.Errorf("Expected result.Version for '%s' to be '20.0.1387.64', but got '%s'", `Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Safari/537.36 OPR/20.0.1387.64`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`, result.Category)
		}
		if result.Name != "Internet Explorer" {
			t.Errorf("Expected result.Name for '%s' to be 'Internet Explorer', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`, result.OsVersion)
		}
		if result.Version != "6.0" {
			t.Errorf("Expected result.Version for '%s' to be '6.0', but got '%s'", `Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1) Sleipnir/2.8.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`, result.Name)
		}
		if result.Os != "Windows Vista" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows Vista', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`, result.Os)
		}
		if result.OsVersion != "NT 6.0" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.0', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows; U; Windows NT 6.0; ja-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`, result.Category)
		}
		if result.Name != "UNKNOWN" {
			t.Errorf("Expected result.Name for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`, result.Name)
		}
		if result.Os != "Windows XP" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows XP', but got '%s'", `Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`, result.Os)
		}
		if result.OsVersion != "NT 5.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 5.1', but got '%s'", `Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Windows NT 5.1; rv:8.0) Gecko/20111105 Thunderbird/8.0`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36': %s`, err)
	} else {
		if result.Category != "pc" {
			t.Errorf("Expected result.Category for '%s' to be 'pc', but got '%s'", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`, result.Category)
		}
		if result.Name != "Yandex Browser" {
			t.Errorf("Expected result.Name for '%s' to be 'Yandex Browser', but got '%s'", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`, result.Name)
		}
		if result.Os != "Windows 7" {
			t.Errorf("Expected result.Os for '%s' to be 'Windows 7', but got '%s'", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`, result.Os)
		}
		if result.OsVersion != "NT 6.1" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'NT 6.1', but got '%s'", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`, result.OsVersion)
		}
		if result.Version != "18.1.1.839" {
			t.Errorf("Expected result.Version for '%s' to be '18.1.1.839', but got '%s'", `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 YaBrowser/18.1.1.839 Yowser/2.5 Safari/537.36`, result.Version)
		}
	}

}
