/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_misc(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`AppleSyndication/56.1`)
	if err != nil {
		t.Errorf(`Failed to parse 'AppleSyndication/56.1': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `AppleSyndication/56.1`, result.Category)
		}
		if result.Name != "Safari RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'Safari RSSReader', but got '%s'", `AppleSyndication/56.1`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `AppleSyndication/56.1`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `AppleSyndication/56.1`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `AppleSyndication/56.1`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`, result.Category)
		}
		if result.Name != "Google Desktop" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Desktop', but got '%s'", `Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`, result.Version)
		}
	}
	result, err = Parse(`Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`, result.Category)
		}
		if result.Name != "Windows RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'Windows RSSReader', but got '%s'", `Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`, result.Version)
		}
	}
	result, err = Parse(`RssBar/1.29 (RssBar for unDonut 1.35)`)
	if err != nil {
		t.Errorf(`Failed to parse 'RssBar/1.29 (RssBar for unDonut 1.35)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `RssBar/1.29 (RssBar for unDonut 1.35)`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `RssBar/1.29 (RssBar for unDonut 1.35)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `RssBar/1.29 (RssBar for unDonut 1.35)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `RssBar/1.29 (RssBar for unDonut 1.35)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `RssBar/1.29 (RssBar for unDonut 1.35)`, result.Version)
		}
	}
	result, err = Parse(`MagpieRSS/0.61 (+http://magpierss.sf.net)`)
	if err != nil {
		t.Errorf(`Failed to parse 'MagpieRSS/0.61 (+http://magpierss.sf.net)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `MagpieRSS/0.61 (+http://magpierss.sf.net)`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `MagpieRSS/0.61 (+http://magpierss.sf.net)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `MagpieRSS/0.61 (+http://magpierss.sf.net)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `MagpieRSS/0.61 (+http://magpierss.sf.net)`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `MagpieRSS/0.61 (+http://magpierss.sf.net)`, result.Version)
		}
	}
	result, err = Parse(`gooRSSreader3.7 - build20090410`)
	if err != nil {
		t.Errorf(`Failed to parse 'gooRSSreader3.7 - build20090410': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `gooRSSreader3.7 - build20090410`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `gooRSSreader3.7 - build20090410`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `gooRSSreader3.7 - build20090410`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `gooRSSreader3.7 - build20090410`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `gooRSSreader3.7 - build20090410`, result.Version)
		}
	}
	result, err = Parse(`Fenrir Headline-Reader Plugin`)
	if err != nil {
		t.Errorf(`Failed to parse 'Fenrir Headline-Reader Plugin': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Fenrir Headline-Reader Plugin`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `Fenrir Headline-Reader Plugin`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Fenrir Headline-Reader Plugin`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Fenrir Headline-Reader Plugin`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Fenrir Headline-Reader Plugin`, result.Version)
		}
	}
	result, err = Parse(`jsRSS++/3.15`)
	if err != nil {
		t.Errorf(`Failed to parse 'jsRSS++/3.15': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `jsRSS++/3.15`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `jsRSS++/3.15`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `jsRSS++/3.15`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `jsRSS++/3.15`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `jsRSS++/3.15`, result.Version)
		}
	}
	result, err = Parse(`cococ/1.06`)
	if err != nil {
		t.Errorf(`Failed to parse 'cococ/1.06': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `cococ/1.06`, result.Category)
		}
		if result.Name != "RSSReader" {
			t.Errorf("Expected result.Name for '%s' to be 'RSSReader', but got '%s'", `cococ/1.06`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `cococ/1.06`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `cococ/1.06`, result.OsVersion)
		}
		if result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `cococ/1.06`, result.Version)
		}
	}
	result, err = Parse(`Wget/1.12 (linux-gnu)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Wget/1.12 (linux-gnu)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Wget/1.12 (linux-gnu)`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Wget/1.12 (linux-gnu)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Wget/1.12 (linux-gnu)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Wget/1.12 (linux-gnu)`, result.OsVersion)
		}
		if result.Version != "wget" {
			t.Errorf("Expected result.Version for '%s' to be 'wget', but got '%s'", `Wget/1.12 (linux-gnu)`, result.Version)
		}
	}
	result, err = Parse(`Apache-HttpClient/UNAVAILABLE (java 1.4)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Apache-HttpClient/UNAVAILABLE (java 1.4)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Apache-HttpClient/UNAVAILABLE (java 1.4)`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Apache-HttpClient/UNAVAILABLE (java 1.4)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Apache-HttpClient/UNAVAILABLE (java 1.4)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Apache-HttpClient/UNAVAILABLE (java 1.4)`, result.OsVersion)
		}
		if result.Version != "Java" {
			t.Errorf("Expected result.Version for '%s' to be 'Java', but got '%s'", `Apache-HttpClient/UNAVAILABLE (java 1.4)`, result.Version)
		}
	}
	result, err = Parse(`livedoor HttpClient`)
	if err != nil {
		t.Errorf(`Failed to parse 'livedoor HttpClient': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `livedoor HttpClient`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `livedoor HttpClient`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `livedoor HttpClient`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `livedoor HttpClient`, result.OsVersion)
		}
		if result.Version != "Java" {
			t.Errorf("Expected result.Version for '%s' to be 'Java', but got '%s'", `livedoor HttpClient`, result.Version)
		}
	}
	result, err = Parse(`Jakarta Commons-HttpClient/3.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Jakarta Commons-HttpClient/3.0': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Jakarta Commons-HttpClient/3.0`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Jakarta Commons-HttpClient/3.0`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Jakarta Commons-HttpClient/3.0`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Jakarta Commons-HttpClient/3.0`, result.OsVersion)
		}
		if result.Version != "Java" {
			t.Errorf("Expected result.Version for '%s' to be 'Java', but got '%s'", `Jakarta Commons-HttpClient/3.0`, result.Version)
		}
	}
	result, err = Parse(`Java/1.5.0_17`)
	if err != nil {
		t.Errorf(`Failed to parse 'Java/1.5.0_17': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Java/1.5.0_17`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Java/1.5.0_17`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Java/1.5.0_17`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Java/1.5.0_17`, result.OsVersion)
		}
		if result.Version != "Java" {
			t.Errorf("Expected result.Version for '%s' to be 'Java', but got '%s'", `Java/1.5.0_17`, result.Version)
		}
	}
	result, err = Parse(`IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`)
	if err != nil {
		t.Errorf(`Failed to parse 'IE6.0,Java(TM) 2 Runtime Environment, Standard Edition': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`, result.OsVersion)
		}
		if result.Version != "Java" {
			t.Errorf("Expected result.Version for '%s' to be 'Java', but got '%s'", `IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`, result.Version)
		}
	}
	result, err = Parse(`libwww-perl/5.835`)
	if err != nil {
		t.Errorf(`Failed to parse 'libwww-perl/5.835': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `libwww-perl/5.835`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `libwww-perl/5.835`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `libwww-perl/5.835`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `libwww-perl/5.835`, result.OsVersion)
		}
		if result.Version != "perl" {
			t.Errorf("Expected result.Version for '%s' to be 'perl', but got '%s'", `libwww-perl/5.835`, result.Version)
		}
	}
	result, err = Parse(`WWW-Mechanize/1.64`)
	if err != nil {
		t.Errorf(`Failed to parse 'WWW-Mechanize/1.64': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `WWW-Mechanize/1.64`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `WWW-Mechanize/1.64`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `WWW-Mechanize/1.64`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `WWW-Mechanize/1.64`, result.OsVersion)
		}
		if result.Version != "perl" {
			t.Errorf("Expected result.Version for '%s' to be 'perl', but got '%s'", `WWW-Mechanize/1.64`, result.Version)
		}
	}
	result, err = Parse(`LWP::Simple/5.800`)
	if err != nil {
		t.Errorf(`Failed to parse 'LWP::Simple/5.800': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `LWP::Simple/5.800`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `LWP::Simple/5.800`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `LWP::Simple/5.800`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `LWP::Simple/5.800`, result.OsVersion)
		}
		if result.Version != "perl" {
			t.Errorf("Expected result.Version for '%s' to be 'perl', but got '%s'", `LWP::Simple/5.800`, result.Version)
		}
	}
	result, err = Parse(`LWP LDMusicNews::LDNewsAPI`)
	if err != nil {
		t.Errorf(`Failed to parse 'LWP LDMusicNews::LDNewsAPI': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `LWP LDMusicNews::LDNewsAPI`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `LWP LDMusicNews::LDNewsAPI`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `LWP LDMusicNews::LDNewsAPI`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `LWP LDMusicNews::LDNewsAPI`, result.OsVersion)
		}
		if result.Version != "perl" {
			t.Errorf("Expected result.Version for '%s' to be 'perl', but got '%s'", `LWP LDMusicNews::LDNewsAPI`, result.Version)
		}
	}
	result, err = Parse(`lwp-trivial/1.41`)
	if err != nil {
		t.Errorf(`Failed to parse 'lwp-trivial/1.41': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `lwp-trivial/1.41`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `lwp-trivial/1.41`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `lwp-trivial/1.41`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `lwp-trivial/1.41`, result.OsVersion)
		}
		if result.Version != "perl" {
			t.Errorf("Expected result.Version for '%s' to be 'perl', but got '%s'", `lwp-trivial/1.41`, result.Version)
		}
	}
	result, err = Parse(`Ruby`)
	if err != nil {
		t.Errorf(`Failed to parse 'Ruby': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Ruby`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Ruby`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Ruby`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Ruby`, result.OsVersion)
		}
		if result.Version != "ruby" {
			t.Errorf("Expected result.Version for '%s' to be 'ruby', but got '%s'", `Ruby`, result.Version)
		}
	}
	result, err = Parse(`feedzirra http://github.com/pauldix/feedzirra/tree/master`)
	if err != nil {
		t.Errorf(`Failed to parse 'feedzirra http://github.com/pauldix/feedzirra/tree/master': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `feedzirra http://github.com/pauldix/feedzirra/tree/master`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `feedzirra http://github.com/pauldix/feedzirra/tree/master`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `feedzirra http://github.com/pauldix/feedzirra/tree/master`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `feedzirra http://github.com/pauldix/feedzirra/tree/master`, result.OsVersion)
		}
		if result.Version != "ruby" {
			t.Errorf("Expected result.Version for '%s' to be 'ruby', but got '%s'", `feedzirra http://github.com/pauldix/feedzirra/tree/master`, result.Version)
		}
	}
	result, err = Parse(`Typhoeus - https://github.com/typhoeus/typhoeus`)
	if err != nil {
		t.Errorf(`Failed to parse 'Typhoeus - https://github.com/typhoeus/typhoeus': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Typhoeus - https://github.com/typhoeus/typhoeus`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Typhoeus - https://github.com/typhoeus/typhoeus`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Typhoeus - https://github.com/typhoeus/typhoeus`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Typhoeus - https://github.com/typhoeus/typhoeus`, result.OsVersion)
		}
		if result.Version != "ruby" {
			t.Errorf("Expected result.Version for '%s' to be 'ruby', but got '%s'", `Typhoeus - https://github.com/typhoeus/typhoeus`, result.Version)
		}
	}
	result, err = Parse(`Python-urllib/1.16`)
	if err != nil {
		t.Errorf(`Failed to parse 'Python-urllib/1.16': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Python-urllib/1.16`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Python-urllib/1.16`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Python-urllib/1.16`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Python-urllib/1.16`, result.OsVersion)
		}
		if result.Version != "python" {
			t.Errorf("Expected result.Version for '%s' to be 'python', but got '%s'", `Python-urllib/1.16`, result.Version)
		}
	}
	result, err = Parse(`Twisted PageGetter`)
	if err != nil {
		t.Errorf(`Failed to parse 'Twisted PageGetter': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Twisted PageGetter`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Twisted PageGetter`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Twisted PageGetter`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Twisted PageGetter`, result.OsVersion)
		}
		if result.Version != "python" {
			t.Errorf("Expected result.Version for '%s' to be 'python', but got '%s'", `Twisted PageGetter`, result.Version)
		}
	}
	result, err = Parse(`PHP/5.2.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'PHP/5.2.13': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `PHP/5.2.13`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `PHP/5.2.13`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `PHP/5.2.13`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `PHP/5.2.13`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `PHP/5.2.13`, result.Version)
		}
	}
	result, err = Parse(`PHP`)
	if err != nil {
		t.Errorf(`Failed to parse 'PHP': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `PHP`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `PHP`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `PHP`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `PHP`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `PHP`, result.Version)
		}
	}
	result, err = Parse(`PEAR HTTP_Request class`)
	if err != nil {
		t.Errorf(`Failed to parse 'PEAR HTTP_Request class': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `PEAR HTTP_Request class`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `PEAR HTTP_Request class`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `PEAR HTTP_Request class`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `PEAR HTTP_Request class`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `PEAR HTTP_Request class`, result.Version)
		}
	}
	result, err = Parse(`HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`)
	if err != nil {
		t.Errorf(`Failed to parse 'HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`, result.Version)
		}
	}
	result, err = Parse(`PECL::HTTP/1.7.4 (PHP/5.4.4)`)
	if err != nil {
		t.Errorf(`Failed to parse 'PECL::HTTP/1.7.4 (PHP/5.4.4)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `PECL::HTTP/1.7.4 (PHP/5.4.4)`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `PECL::HTTP/1.7.4 (PHP/5.4.4)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `PECL::HTTP/1.7.4 (PHP/5.4.4)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `PECL::HTTP/1.7.4 (PHP/5.4.4)`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `PECL::HTTP/1.7.4 (PHP/5.4.4)`, result.Version)
		}
	}
	result, err = Parse(`WordPress/3.2.1; http://www.painlog.jp`)
	if err != nil {
		t.Errorf(`Failed to parse 'WordPress/3.2.1; http://www.painlog.jp': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `WordPress/3.2.1; http://www.painlog.jp`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `WordPress/3.2.1; http://www.painlog.jp`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `WordPress/3.2.1; http://www.painlog.jp`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `WordPress/3.2.1; http://www.painlog.jp`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `WordPress/3.2.1; http://www.painlog.jp`, result.Version)
		}
	}
	result, err = Parse(`CakePHP`)
	if err != nil {
		t.Errorf(`Failed to parse 'CakePHP': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `CakePHP`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `CakePHP`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `CakePHP`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `CakePHP`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `CakePHP`, result.Version)
		}
	}
	result, err = Parse(`PukiWiki/1.4.6`)
	if err != nil {
		t.Errorf(`Failed to parse 'PukiWiki/1.4.6': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `PukiWiki/1.4.6`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `PukiWiki/1.4.6`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `PukiWiki/1.4.6`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `PukiWiki/1.4.6`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `PukiWiki/1.4.6`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; PEAR HTTP_Request class;)': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`, result.OsVersion)
		}
		if result.Version != "php" {
			t.Errorf("Expected result.Version for '%s' to be 'php', but got '%s'", `Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`, result.Version)
		}
	}
	result, err = Parse(`curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`)
	if err != nil {
		t.Errorf(`Failed to parse 'curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2': %s`, err)
	} else {
		if result.Category != "misc" {
			t.Errorf("Expected result.Category for '%s' to be 'misc', but got '%s'", `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`, result.Category)
		}
		if result.Name != "HTTP Library" {
			t.Errorf("Expected result.Name for '%s' to be 'HTTP Library', but got '%s'", `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`, result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`, result.Os)
		}
		if result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`, result.OsVersion)
		}
		if result.Version != "curl" {
			t.Errorf("Expected result.Version for '%s' to be 'curl', but got '%s'", `curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.19.1 Basic ECC zlib/1.2.3 libidn/1.18 libssh2/1.4.2`, result.Version)
		}
	}

}
