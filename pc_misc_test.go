/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_pc_misc(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_4; ja-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.20.1': %s`, err)
    } else {
        if result.Name != "Safari" {
            t.Errorf("Expected result.Name to be 'Safari', but got '%s'", result.Name)
        }
        if result.Version != "3.1.2" {
            t.Errorf("Expected result.Version to be '3.1.2', but got '%s'", result.Version)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36': %s`, err)
    } else {
        if result.Version != "27.0.1453.93" {
            t.Errorf("Expected result.Version to be '27.0.1453.93', but got '%s'", result.Version)
        }
        if result.Name != "Chrome" {
            t.Errorf("Expected result.Name to be 'Chrome', but got '%s'", result.Name)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:21.0) Gecko/20100101 Firefox/21.0': %s`, err)
    } else {
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Version != "21.0" {
            t.Errorf("Expected result.Version to be '21.0', but got '%s'", result.Version)
        }
        if result.Name != "Firefox" {
            t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
        }
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15`)
    if err != nil {
        t.Errorf(`Failed to parse 'Opera/9.80 (Macintosh; Intel Mac OS X 10.8.3) Presto/2.12.388 Version/12.15': %s`, err)
    } else {
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
        if result.Name != "Opera" {
            t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
        }
        if result.Version != "12.15" {
            t.Errorf("Expected result.Version to be '12.15', but got '%s'", result.Version)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Instapaper/4.0 (+http://www.instapaper.com/)': %s`, err)
    } else {
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; ja-JP-mac; rv:1.9.1.19) Gecko/20110420 SeaMonkey/2.0.14': %s`, err)
    } else {
        if result.Os != "Mac OSX" {
            t.Errorf("Expected result.Os to be 'Mac OSX', but got '%s'", result.Os)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Macintosh; U; PPC; en-US; mimic; rv:9.2.1) (mimic Gecko/20100722 Firefox/3.6.8) Classilla/CFM': %s`, err)
    } else {
        if result.Os != "Mac OS Classic" {
            t.Errorf("Expected result.Os to be 'Mac OS Classic', but got '%s'", result.Os)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Version != "3.6.8" {
            t.Errorf("Expected result.Version to be '3.6.8', but got '%s'", result.Version)
        }
        if result.Name != "Firefox" {
            t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)': %s`, err)
    } else {
        if result.Name != "Internet Explorer" {
            t.Errorf("Expected result.Name to be 'Internet Explorer', but got '%s'", result.Name)
        }
        if result.Version != "5.17" {
            t.Errorf("Expected result.Version to be '5.17', but got '%s'", result.Version)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Os != "Mac OS Classic" {
            t.Errorf("Expected result.Os to be 'Mac OS Classic', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Ubuntu; X11; Linux i686; rv:9.0.1) Gecko/20100101 Firefox/9.0.1': %s`, err)
    } else {
        if result.Os != "Linux" {
            t.Errorf("Expected result.Os to be 'Linux', but got '%s'", result.Os)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "Firefox" {
            t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
        }
        if result.Version != "9.0.1" {
            t.Errorf("Expected result.Version to be '9.0.1', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7': %s`, err)
    } else {
        if result.Os != "Linux" {
            t.Errorf("Expected result.Os to be 'Linux', but got '%s'", result.Os)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "Chrome" {
            t.Errorf("Expected result.Name to be 'Chrome', but got '%s'", result.Name)
        }
        if result.Version != "16.0.912.75" {
            t.Errorf("Expected result.Version to be '16.0.912.75', but got '%s'", result.Version)
        }
    }
    result, err = Parse(`Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.24) Gecko cXense': %s`, err)
    } else {
        if result.Os != "Linux" {
            t.Errorf("Expected result.Os to be 'Linux', but got '%s'", result.Os)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux i686; ja-JP; rv:1.8.1.23) Gecko/20090910 SeaMonkey/1.1.18': %s`, err)
    } else {
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
        if result.Os != "Linux" {
            t.Errorf("Expected result.Os to be 'Linux', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 Iceweasel/3.0.3 (Debian-3.0.3-3)': %s`, err)
    } else {
        if result.Os != "Linux" {
            t.Errorf("Expected result.Os to be 'Linux', but got '%s'", result.Os)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (X11; FreeBSD amd64; rv:8.0) Gecko/20100101 Firefox/8.0': %s`, err)
    } else {
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Version != "8.0" {
            t.Errorf("Expected result.Version to be '8.0', but got '%s'", result.Version)
        }
        if result.Name != "Firefox" {
            t.Errorf("Expected result.Name to be 'Firefox', but got '%s'", result.Name)
        }
        if result.Os != "BSD" {
            t.Errorf("Expected result.Os to be 'BSD', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52`)
    if err != nil {
        t.Errorf(`Failed to parse 'Opera/9.80 (X11; FreeBSD 8.2-RELEASE-p3 amd64; U; ja) Presto/2.9.168 Version/11.52': %s`, err)
    } else {
        if result.Os != "BSD" {
            t.Errorf("Expected result.Os to be 'BSD', but got '%s'", result.Os)
        }
        if result.Version != "11.52" {
            t.Errorf("Expected result.Version to be '11.52', but got '%s'", result.Version)
        }
        if result.Name != "Opera" {
            t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
        }
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
    }
}
