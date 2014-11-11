/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_mobilephone_docomo(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`DoCoMo/2.0 SH01A(c100;TB;W24H16)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 SH01A(c100;TB;W24H16)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `DoCoMo/2.0 SH01A(c100;TB;W24H16)`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `DoCoMo/2.0 SH01A(c100;TB;W24H16)`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `DoCoMo/2.0 SH01A(c100;TB;W24H16)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 SH01A(c100;TB;W24H16)`, result.OsVersion)
		}
		if result.Version != "SH01A" {
			t.Errorf("Expected result.Version for '%s' to be 'SH01A', but got '%s'", `DoCoMo/2.0 SH01A(c100;TB;W24H16)`, result.Version)
		}
	}
	result, err = Parse(`DoCoMo/2.0 N07B(c500;TB;W24H16)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 N07B(c500;TB;W24H16)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `DoCoMo/2.0 N07B(c500;TB;W24H16)`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `DoCoMo/2.0 N07B(c500;TB;W24H16)`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `DoCoMo/2.0 N07B(c500;TB;W24H16)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 N07B(c500;TB;W24H16)`, result.OsVersion)
		}
		if result.Version != "N07B" {
			t.Errorf("Expected result.Version for '%s' to be 'N07B', but got '%s'", `DoCoMo/2.0 N07B(c500;TB;W24H16)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`, result.OsVersion)
		}
		if result.Version != "D505i" {
			t.Errorf("Expected result.Version for '%s' to be 'D505i', but got '%s'", `Mozilla/5.0 (compatible; DoCoMo/1.0/D505i/c20/TB/W20H10; http://www.rcdtokyo.com/pc2m/)`, result.Version)
		}
	}
	result, err = Parse(`DoCoMo/1.0/N505i/c20/TB/W24H12`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/1.0/N505i/c20/TB/W24H12': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `DoCoMo/1.0/N505i/c20/TB/W24H12`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `DoCoMo/1.0/N505i/c20/TB/W24H12`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `DoCoMo/1.0/N505i/c20/TB/W24H12`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/1.0/N505i/c20/TB/W24H12`, result.OsVersion)
		}
		if result.Version != "N505i" {
			t.Errorf("Expected result.Version for '%s' to be 'N505i', but got '%s'", `DoCoMo/1.0/N505i/c20/TB/W24H12`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`, result.OsVersion)
		}
		if result.Version != "N905imyu_W" {
			t.Errorf("Expected result.Version for '%s' to be 'N905imyu_W', but got '%s'", `Mozilla/4.08 (N905imyu_W;FOMA;c500;TB)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (F02B;FOMA;like Gecko)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (F02B;FOMA;like Gecko)': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Mozilla/5.0 (F02B;FOMA;like Gecko)`, result.Category)
		}
		if result.Name != "docomo" {
			t.Errorf("Expected result.Name for '%s' to be 'docomo', but got '%s'", `Mozilla/5.0 (F02B;FOMA;like Gecko)`, result.Name)
		}
		if result.Os != "docomo" {
			t.Errorf("Expected result.Os for '%s' to be 'docomo', but got '%s'", `Mozilla/5.0 (F02B;FOMA;like Gecko)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (F02B;FOMA;like Gecko)`, result.OsVersion)
		}
		if result.Version != "F02B" {
			t.Errorf("Expected result.Version for '%s' to be 'F02B', but got '%s'", `Mozilla/5.0 (F02B;FOMA;like Gecko)`, result.Version)
		}
	}

}
