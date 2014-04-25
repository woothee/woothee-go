/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_misc(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`AppleSyndication/56.1`)
    if err != nil {
        t.Errorf(`Failed to parse 'AppleSyndication/56.1': %s`, err)
    } else {
        if result.Name != "Safari RSSReader" {
            t.Errorf("Expected result.Name to be 'Safari RSSReader', but got '%s'", result.Name)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Google Desktop/5.9.1005.12335; http://desktop.google.com/)': %s`, err)
    } else {
        if result.Name != "Google Desktop" {
            t.Errorf("Expected result.Name to be 'Google Desktop', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Windows-RSS-Platform/2.0 (MSIE 9.0; Windows NT 6.0)': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Name != "Windows RSSReader" {
            t.Errorf("Expected result.Name to be 'Windows RSSReader', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`RssBar/1.29 (RssBar for unDonut 1.35)`)
    if err != nil {
        t.Errorf(`Failed to parse 'RssBar/1.29 (RssBar for unDonut 1.35)': %s`, err)
    } else {
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`MagpieRSS/0.61 (+http://magpierss.sf.net)`)
    if err != nil {
        t.Errorf(`Failed to parse 'MagpieRSS/0.61 (+http://magpierss.sf.net)': %s`, err)
    } else {
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`gooRSSreader3.7 - build20090410`)
    if err != nil {
        t.Errorf(`Failed to parse 'gooRSSreader3.7 - build20090410': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Fenrir Headline-Reader Plugin`)
    if err != nil {
        t.Errorf(`Failed to parse 'Fenrir Headline-Reader Plugin': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`jsRSS++/3.15`)
    if err != nil {
        t.Errorf(`Failed to parse 'jsRSS++/3.15': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`cococ/1.06`)
    if err != nil {
        t.Errorf(`Failed to parse 'cococ/1.06': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "RSSReader" {
            t.Errorf("Expected result.Name to be 'RSSReader', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Wget/1.12 (linux-gnu)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Wget/1.12 (linux-gnu)': %s`, err)
    } else {
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "wget" {
            t.Errorf("Expected result.Version to be 'wget', but got '%s'", result.Version)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Apache-HttpClient/UNAVAILABLE (java 1.4)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Apache-HttpClient/UNAVAILABLE (java 1.4)': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "Java" {
            t.Errorf("Expected result.Version to be 'Java', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`livedoor HttpClient`)
    if err != nil {
        t.Errorf(`Failed to parse 'livedoor HttpClient': %s`, err)
    } else {
        if result.Version != "Java" {
            t.Errorf("Expected result.Version to be 'Java', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Jakarta Commons-HttpClient/3.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Jakarta Commons-HttpClient/3.0': %s`, err)
    } else {
        if result.Version != "Java" {
            t.Errorf("Expected result.Version to be 'Java', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Java/1.5.0_17`)
    if err != nil {
        t.Errorf(`Failed to parse 'Java/1.5.0_17': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Version != "Java" {
            t.Errorf("Expected result.Version to be 'Java', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`IE6.0,Java(TM) 2 Runtime Environment, Standard Edition`)
    if err != nil {
        t.Errorf(`Failed to parse 'IE6.0,Java(TM) 2 Runtime Environment, Standard Edition': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Version != "Java" {
            t.Errorf("Expected result.Version to be 'Java', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`libwww-perl/5.835`)
    if err != nil {
        t.Errorf(`Failed to parse 'libwww-perl/5.835': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "perl" {
            t.Errorf("Expected result.Version to be 'perl', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`WWW-Mechanize/1.64`)
    if err != nil {
        t.Errorf(`Failed to parse 'WWW-Mechanize/1.64': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "perl" {
            t.Errorf("Expected result.Version to be 'perl', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`LWP::Simple/5.800`)
    if err != nil {
        t.Errorf(`Failed to parse 'LWP::Simple/5.800': %s`, err)
    } else {
        if result.Version != "perl" {
            t.Errorf("Expected result.Version to be 'perl', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`LWP LDMusicNews::LDNewsAPI`)
    if err != nil {
        t.Errorf(`Failed to parse 'LWP LDMusicNews::LDNewsAPI': %s`, err)
    } else {
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "perl" {
            t.Errorf("Expected result.Version to be 'perl', but got '%s'", result.Version)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`lwp-trivial/1.41`)
    if err != nil {
        t.Errorf(`Failed to parse 'lwp-trivial/1.41': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "perl" {
            t.Errorf("Expected result.Version to be 'perl', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Ruby`)
    if err != nil {
        t.Errorf(`Failed to parse 'Ruby': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Version != "ruby" {
            t.Errorf("Expected result.Version to be 'ruby', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`feedzirra http://github.com/pauldix/feedzirra/tree/master`)
    if err != nil {
        t.Errorf(`Failed to parse 'feedzirra http://github.com/pauldix/feedzirra/tree/master': %s`, err)
    } else {
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "ruby" {
            t.Errorf("Expected result.Version to be 'ruby', but got '%s'", result.Version)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Typhoeus - https://github.com/typhoeus/typhoeus`)
    if err != nil {
        t.Errorf(`Failed to parse 'Typhoeus - https://github.com/typhoeus/typhoeus': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "ruby" {
            t.Errorf("Expected result.Version to be 'ruby', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Python-urllib/1.16`)
    if err != nil {
        t.Errorf(`Failed to parse 'Python-urllib/1.16': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "python" {
            t.Errorf("Expected result.Version to be 'python', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Twisted PageGetter`)
    if err != nil {
        t.Errorf(`Failed to parse 'Twisted PageGetter': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "python" {
            t.Errorf("Expected result.Version to be 'python', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`PHP/5.2.13`)
    if err != nil {
        t.Errorf(`Failed to parse 'PHP/5.2.13': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`PHP`)
    if err != nil {
        t.Errorf(`Failed to parse 'PHP': %s`, err)
    } else {
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`PEAR HTTP_Request class`)
    if err != nil {
        t.Errorf(`Failed to parse 'PEAR HTTP_Request class': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6`)
    if err != nil {
        t.Errorf(`Failed to parse 'HTTP_Request2/2.1.1 (http://pear.php.net/package/http_request2) PHP/5.3.10-1ubuntu3.6': %s`, err)
    } else {
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`PECL::HTTP/1.7.4 (PHP/5.4.4)`)
    if err != nil {
        t.Errorf(`Failed to parse 'PECL::HTTP/1.7.4 (PHP/5.4.4)': %s`, err)
    } else {
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`WordPress/3.2.1; http://www.painlog.jp`)
    if err != nil {
        t.Errorf(`Failed to parse 'WordPress/3.2.1; http://www.painlog.jp': %s`, err)
    } else {
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`CakePHP`)
    if err != nil {
        t.Errorf(`Failed to parse 'CakePHP': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`PukiWiki/1.4.6`)
    if err != nil {
        t.Errorf(`Failed to parse 'PukiWiki/1.4.6': %s`, err)
    } else {
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; PEAR HTTP_Request class;)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; PEAR HTTP_Request class;)': %s`, err)
    } else {
        if result.Category != "misc" {
            t.Errorf("Expected result.Category to be 'misc', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Version != "php" {
            t.Errorf("Expected result.Version to be 'php', but got '%s'", result.Version)
        }
        if result.Name != "HTTP Library" {
            t.Errorf("Expected result.Name to be 'HTTP Library', but got '%s'", result.Name)
        }
    }
}
