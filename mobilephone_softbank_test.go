/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_mobilephone_softbank(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.OsVersion)
		}
		if result.Version != "841SH" {
			t.Errorf("Expected result.Version for '%s' to be '841SH', but got '%s'", `SoftBank/1.0/841SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Version)
		}
	}
	result, err = Parse(`SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.OsVersion)
		}
		if result.Version != "004SH" {
			t.Errorf("Expected result.Version for '%s' to be '004SH', but got '%s'", `SoftBank/2.0/004SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Version)
		}
	}
	result, err = Parse(`SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.OsVersion)
		}
		if result.Version != "944SH" {
			t.Errorf("Expected result.Version for '%s' to be '944SH', but got '%s'", `SoftBank/2.0/944SH/SHJ001/SN000000000000000 Browser/NetFront/3.5 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Version)
		}
	}
	result, err = Parse(`SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`, result.OsVersion)
		}
		if result.Version != "821SC" {
			t.Errorf("Expected result.Version for '%s' to be '821SC', but got '%s'", `SoftBank/1.0/821SC/SCJ001/SN000000000000000 Flash/Flash-Lite/2.0`, result.Version)
		}
	}
	result, err = Parse(`Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.OsVersion)
		}
		if result.Version != "V905SH" {
			t.Errorf("Expected result.Version for '%s' to be 'V905SH', but got '%s'", `Vodafone/1.0/V905SH/SHJ001 Browser/VF-NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`, result.Version)
		}
	}
	result, err = Parse(`Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`, result.OsVersion)
		}
		if result.Version != "V804SH" {
			t.Errorf("Expected result.Version for '%s' to be 'V804SH', but got '%s'", `Vodafone/1.0/V804SH/SHJ001/SN0000000000000 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.3.2 Ext-V-Profile/VSCL-2.0.0`, result.Version)
		}
	}
	result, err = Parse(`J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0': %s`, err)
	} else {
		if result.Category != "mobilephone" {
			t.Errorf("Expected result.Category for '%s' to be 'mobilephone', but got '%s'", `J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`, result.Category)
		}
		if result.Name != "SoftBank Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'SoftBank Mobile', but got '%s'", `J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`, result.Name)
		}
		if result.Os != "SoftBank" {
			t.Errorf("Expected result.Os for '%s' to be 'SoftBank', but got '%s'", `J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`, result.OsVersion)
		}
		if result.Version != "J-T51" {
			t.Errorf("Expected result.Version for '%s' to be 'J-T51', but got '%s'", `J-PHONE/4.0/J-T51/SNXXXXXXXXX000000000 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`, result.Version)
		}
	}

}
