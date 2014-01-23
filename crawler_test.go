/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_crawler(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Yahoo! Slurp" {
            t.Errorf("Expected result.Name to be 'Yahoo! Slurp', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Yahoo! Slurp" {
            t.Errorf("Expected result.Name to be 'Yahoo! Slurp', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Y!J-BRO/YFSJ crawler (compatible; Mozilla 4.0; MSIE 5.5; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html; YahooFeedSeekerJp/2.0; users 0; views 248)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Y!J-BRO/YFSJ crawler (compatible; Mozilla 4.0; MSIE 5.5; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html; YahooFeedSeekerJp/2.0; users 0; views 248)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Yahoo! Japan" {
            t.Errorf("Expected result.Name to be 'Yahoo! Japan', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Y!J-BRP/YFSBJ crawler (compatible; Mozilla 4.0; MSIE 5.5; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html; YahooFeedSeekerBetaJp/2.0; users 0; views 80)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Y!J-BRP/YFSBJ crawler (compatible; Mozilla 4.0; MSIE 5.5; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html; YahooFeedSeekerBetaJp/2.0; users 0; views 80)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Yahoo! Japan" {
            t.Errorf("Expected result.Name to be 'Yahoo! Japan', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Y!J-BRJ/YATS crawler (http://listing.yahoo.co.jp/support/faq/int/other/other_001.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Y!J-BRJ/YATS crawler (http://listing.yahoo.co.jp/support/faq/int/other/other_001.html)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Yahoo! Japan" {
            t.Errorf("Expected result.Name to be 'Yahoo! Japan', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Yahoo! Japan" {
            t.Errorf("Expected result.Name to be 'Yahoo! Japan', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Y!J-BSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/blog-search/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Y!J-BSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/blog-search/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Yahoo! Japan" {
            t.Errorf("Expected result.Name to be 'Yahoo! Japan', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Yahoo Pipes 2.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Yahoo Pipes 2.0': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Yahoo! Pipes" {
            t.Errorf("Expected result.Name to be 'Yahoo! Pipes', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)': %s`, err)
    } else {
        if result.Name != "Baiduspider" {
            t.Errorf("Expected result.Name to be 'Baiduspider', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Baiduspider+(+http://www.baidu.jp/spider/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Baiduspider+(+http://www.baidu.jp/spider/)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Baiduspider" {
            t.Errorf("Expected result.Name to be 'Baiduspider', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Baiduspider-image+(+http://www.baidu.com/search/spider.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Baiduspider-image+(+http://www.baidu.com/search/spider.htm)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Baiduspider" {
            t.Errorf("Expected result.Name to be 'Baiduspider', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`msnbot/1.1 (+http://search.msn.com/msnbot.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse 'msnbot/1.1 (+http://search.msn.com/msnbot.htm)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "msnbot" {
            t.Errorf("Expected result.Name to be 'msnbot', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse 'msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "msnbot" {
            t.Errorf("Expected result.Name to be 'msnbot', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`msnbot/2.0b (+http://search.msn.com/msnbot.htm)._`)
    if err != nil {
        t.Errorf(`Failed to parse 'msnbot/2.0b (+http://search.msn.com/msnbot.htm)._': %s`, err)
    } else {
        if result.Name != "msnbot" {
            t.Errorf("Expected result.Name to be 'msnbot', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`"msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse '"msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "msnbot" {
            t.Errorf("Expected result.Name to be 'msnbot', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse 'msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "msnbot" {
            t.Errorf("Expected result.Name to be 'msnbot', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "bingbot" {
            t.Errorf("Expected result.Name to be 'bingbot', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Yeti/1.0 (NHN Corp.; http://help.naver.com/robots/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Yeti/1.0 (NHN Corp.; http://help.naver.com/robots/)': %s`, err)
    } else {
        if result.Name != "Naver Yeti" {
            t.Errorf("Expected result.Name to be 'Naver Yeti', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/3.0 (compatible; Indy Library)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/3.0 (compatible; Indy Library)': %s`, err)
    } else {
        if result.Name != "Indy Library" {
            t.Errorf("Expected result.Name to be 'Indy Library', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Apple-PubSub/65.28`)
    if err != nil {
        t.Errorf(`Failed to parse 'Apple-PubSub/65.28': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Apple iCloud" {
            t.Errorf("Expected result.Name to be 'Apple iCloud', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`R6_CommentReader(www.radian6.com/crawler)`)
    if err != nil {
        t.Errorf(`Failed to parse 'R6_CommentReader(www.radian6.com/crawler)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "salesforce radian6" {
            t.Errorf("Expected result.Name to be 'salesforce radian6', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`R6_FeedFetcher(www.radian6.com/crawler)`)
    if err != nil {
        t.Errorf(`Failed to parse 'R6_FeedFetcher(www.radian6.com/crawler)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "salesforce radian6" {
            t.Errorf("Expected result.Name to be 'salesforce radian6', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Genieo Web Filter" {
            t.Errorf("Expected result.Name to be 'Genieo Web Filter', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Butterfly/1.0; +http://labs.topsy.com/butterfly/) Gecko/2009032608 Firefox/3.0.8`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Butterfly/1.0; +http://labs.topsy.com/butterfly/) Gecko/2009032608 Firefox/3.0.8': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "topsy Butterfly" {
            t.Errorf("Expected result.Name to be 'topsy Butterfly', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)`)
    if err != nil {
        t.Errorf(`Failed to parse 'rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "SeoMoz rogerbot" {
            t.Errorf("Expected result.Name to be 'SeoMoz rogerbot', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)`)
    if err != nil {
        t.Errorf(`Failed to parse 'rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "SeoMoz rogerbot" {
            t.Errorf("Expected result.Name to be 'SeoMoz rogerbot', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; AhrefsBot/4.0; +http://ahrefs.com/robot/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; AhrefsBot/4.0; +http://ahrefs.com/robot/)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "ahref AhrefsBot" {
            t.Errorf("Expected result.Name to be 'ahref AhrefsBot', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Hatena Antenna/0.5 (http://a.hatena.ne.jp/help)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Hatena Antenna/0.5 (http://a.hatena.ne.jp/help)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Hatena" {
            t.Errorf("Expected result.Name to be 'Hatena', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Hatena Pagetitle Agent/1.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'Hatena Pagetitle Agent/1.0': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "Hatena" {
            t.Errorf("Expected result.Name to be 'Hatena', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`Hatena Diary RSS Module (http://d.hatena.ne.jp/help#rssmodule)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Hatena Diary RSS Module (http://d.hatena.ne.jp/help#rssmodule)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "Hatena" {
            t.Errorf("Expected result.Name to be 'Hatena', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "goo" {
            t.Errorf("Expected result.Name to be 'goo', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "goo" {
            t.Errorf("Expected result.Name to be 'goo', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`gooblogsearch/2.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'gooblogsearch/2.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)': %s`, err)
    } else {
        if result.Name != "goo" {
            t.Errorf("Expected result.Name to be 'goo', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`livedoor FeedFetcher/0.01 (http://reader.livedoor.com/; 999 subscribers)`)
    if err != nil {
        t.Errorf(`Failed to parse 'livedoor FeedFetcher/0.01 (http://reader.livedoor.com/; 999 subscribers)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "livedoor FeedFetcher" {
            t.Errorf("Expected result.Name to be 'livedoor FeedFetcher', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Fastladder FeedFetcher/0.01 (http://fastladder.com/; 27 subscribers)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Fastladder FeedFetcher/0.01 (http://fastladder.com/; 27 subscribers)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "livedoor FeedFetcher" {
            t.Errorf("Expected result.Name to be 'livedoor FeedFetcher', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)`)
    if err != nil {
        t.Errorf(`Failed to parse 'facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "facebook" {
            t.Errorf("Expected result.Name to be 'facebook', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
    }
    result, err = Parse(`mixi-check/1.0 (http://mixi.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'mixi-check/1.0 (http://mixi.jp/)': %s`, err)
    } else {
        if result.Name != "mixi" {
            t.Errorf("Expected result.Name to be 'mixi', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`mixi-news-crawler/1.00 (http://mixi.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'mixi-news-crawler/1.00 (http://mixi.jp/)': %s`, err)
    } else {
        if result.Name != "mixi" {
            t.Errorf("Expected result.Name to be 'mixi', but got '%s'", result.Name)
        }
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`mixi-crawler/2.00 (http://mixi.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'mixi-crawler/2.00 (http://mixi.jp/)': %s`, err)
    } else {
        if result.Os != "UNKNOWN" {
            t.Errorf("Expected result.Os to be 'UNKNOWN', but got '%s'", result.Os)
        }
        if result.Name != "mixi" {
            t.Errorf("Expected result.Name to be 'mixi', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
}
