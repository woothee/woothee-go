/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_crawler_google(t *testing.T) {
	var result *Result
	var err error
	result, err = Parse(`Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`, result.Category)
		}
		if result.Name != "Googlebot" {
			t.Errorf("Expected result.Name for '%s' to be 'Googlebot', but got '%s'", `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`, result.Version)
		}
	}
	result, err = Parse(`Googlebot-Image/1.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Googlebot-Image/1.0': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Googlebot-Image/1.0`, result.Category)
		}
		if result.Name != "Googlebot" {
			t.Errorf("Expected result.Name for '%s' to be 'Googlebot', but got '%s'", `Googlebot-Image/1.0`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Googlebot-Image/1.0`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Googlebot-Image/1.0`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Googlebot-Image/1.0`, result.Version)
		}
	}
	result, err = Parse(`DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Category)
		}
		if result.Name != "Googlebot Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'Googlebot Mobile', but got '%s'", `DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Version)
		}
	}
	result, err = Parse(`SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Category)
		}
		if result.Name != "Googlebot Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'Googlebot Mobile', but got '%s'", `SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`, result.Version)
		}
	}
	result, err = Parse(`DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`, result.Category)
		}
		if result.Name != "Google Mediapartners" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Mediapartners', but got '%s'", `DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`, result.Version)
		}
	}
	result, err = Parse(`Mediapartners-Google`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mediapartners-Google': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Mediapartners-Google`, result.Category)
		}
		if result.Name != "Google Mediapartners" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Mediapartners', but got '%s'", `Mediapartners-Google`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mediapartners-Google`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mediapartners-Google`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mediapartners-Google`, result.Version)
		}
	}
	result, err = Parse(`Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`, result.Category)
		}
		if result.Name != "Google Feedfetcher" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Feedfetcher', but got '%s'", `Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`, result.Version)
		}
	}
	result, err = Parse(`AppEngine-Google`)
	if err != nil {
		t.Errorf(`Failed to parse 'AppEngine-Google': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `AppEngine-Google`, result.Category)
		}
		if result.Name != "Google AppEngine" {
			t.Errorf("Expected result.Name for '%s' to be 'Google AppEngine', but got '%s'", `AppEngine-Google`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `AppEngine-Google`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `AppEngine-Google`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `AppEngine-Google`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`, result.Category)
		}
		if result.Name != "Google Web Preview" {
			t.Errorf("Expected result.Name for '%s' to be 'Google Web Preview', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`, result.Version)
		}
	}
	result, err = Parse(`FeedBurner/1.0 (http://www.FeedBurner.com)`)
	if err != nil {
		t.Errorf(`Failed to parse 'FeedBurner/1.0 (http://www.FeedBurner.com)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `FeedBurner/1.0 (http://www.FeedBurner.com)`, result.Category)
		}
		if result.Name != "Google FeedBurner" {
			t.Errorf("Expected result.Name for '%s' to be 'Google FeedBurner', but got '%s'", `FeedBurner/1.0 (http://www.FeedBurner.com)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `FeedBurner/1.0 (http://www.FeedBurner.com)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `FeedBurner/1.0 (http://www.FeedBurner.com)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `FeedBurner/1.0 (http://www.FeedBurner.com)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`, result.Category)
		}
		if result.Name != "AdsBot-Google-Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'AdsBot-Google-Mobile', but got '%s'", `Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)`, result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`, result.Category)
		}
		if result.Name != "AdsBot-Google-Mobile" {
			t.Errorf("Expected result.Name for '%s' to be 'AdsBot-Google-Mobile', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)`, result.Version)
		}
	}
	result, err = Parse(`AdsBot-Google (+http://www.google.com/adsbot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'AdsBot-Google (+http://www.google.com/adsbot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category for '%s' to be 'crawler', but got '%s'", `AdsBot-Google (+http://www.google.com/adsbot.html)`, result.Category)
		}
		if result.Name != "AdsBot-Google" {
			t.Errorf("Expected result.Name for '%s' to be 'AdsBot-Google', but got '%s'", `AdsBot-Google (+http://www.google.com/adsbot.html)`, result.Name)
		}
		if true && result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os for '%s' to be 'UNKNOWN', but got '%s'", `AdsBot-Google (+http://www.google.com/adsbot.html)`, result.Os)
		}
		if false && result.OsVersion != "UNKNOWN" {
			t.Errorf("Expected result.OsVersion for '%s' to be 'UNKNOWN', but got '%s'", `AdsBot-Google (+http://www.google.com/adsbot.html)`, result.OsVersion)
		}
		if false && result.Version != "UNKNOWN" {
			t.Errorf("Expected result.Version for '%s' to be 'UNKNOWN', but got '%s'", `AdsBot-Google (+http://www.google.com/adsbot.html)`, result.Version)
		}
	}

}
