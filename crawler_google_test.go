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
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Googlebot" {
			t.Errorf("Expected result.Name to be 'Googlebot', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Googlebot-Image/1.0`)
	if err != nil {
		t.Errorf(`Failed to parse 'Googlebot-Image/1.0': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Googlebot" {
			t.Errorf("Expected result.Name to be 'Googlebot', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Googlebot Mobile" {
			t.Errorf("Expected result.Name to be 'Googlebot Mobile', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Googlebot Mobile" {
			t.Errorf("Expected result.Name to be 'Googlebot Mobile', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)`)
	if err != nil {
		t.Errorf(`Failed to parse 'DoCoMo/2.0 SH905i(c100;TB;W24H16) (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google Mediapartners" {
			t.Errorf("Expected result.Name to be 'Google Mediapartners', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mediapartners-Google`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mediapartners-Google': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google Mediapartners" {
			t.Errorf("Expected result.Name to be 'Google Mediapartners', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; feed-id=000000000000000000)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google Feedfetcher" {
			t.Errorf("Expected result.Name to be 'Google Feedfetcher', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`AppEngine-Google`)
	if err != nil {
		t.Errorf(`Failed to parse 'AppEngine-Google': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google AppEngine" {
			t.Errorf("Expected result.Name to be 'Google AppEngine', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (en-us) AppleWebKit/525.13 (KHTML, like Gecko; Google Web Preview) Version/3.1 Safari/525.13': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google Web Preview" {
			t.Errorf("Expected result.Name to be 'Google Web Preview', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`FeedBurner/1.0 (http://www.FeedBurner.com)`)
	if err != nil {
		t.Errorf(`Failed to parse 'FeedBurner/1.0 (http://www.FeedBurner.com)': %s`, err)
	} else {
		if result.Category != "crawler" {
			t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
		}
		if result.Name != "Google FeedBurner" {
			t.Errorf("Expected result.Name to be 'Google FeedBurner', but got '%s'", result.Name)
		}
		if result.Os != "UNKNOWN" {
			t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
		}
	}
}
