/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
    "testing"
)
func Test_crawler_nonmajor(t *testing.T) {
    var result *Result
    var err     error

    result, err = Parse(`emBot-GalaBuzz/Nutch-1.0 (http://emining.jp/; em@galabuzz.jp)`)
    if err != nil {
        t.Errorf(`Failed to parse 'emBot-GalaBuzz/Nutch-1.0 (http://emining.jp/; em@galabuzz.jp)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US; aggregator VocusBot 0.4; +http://www.vocus.com/vnhs.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US; aggregator VocusBot 0.4; +http://www.vocus.com/vnhs.html)': %s`, err)
    } else {
        if result.Category != "pc" {
            t.Errorf("Expected result.Category to be 'pc', but got '%s'", result.Category)
        }
        if result.Os != "Windows Vista" {
            t.Errorf("Expected result.Os to be 'Windows Vista', but got '%s'", result.Os)
        }
        if result.Name != "UNKNOWN" {
            t.Errorf("Expected result.Name to be 'UNKNOWN', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; Rakutenbot/1.0; +http://dynamic.rakuten.co.jp/bot.html)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; Rakutenbot/1.0; +http://dynamic.rakuten.co.jp/bot.html)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`"mapion-news-bot/1.0 (http://www.mapion.co.jp/news/)"`)
    if err != nil {
        t.Errorf(`Failed to parse '"mapion-news-bot/1.0 (http://www.mapion.co.jp/news/)"': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; TweetmemeBot/2.11; +http://tweetmeme.com/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; TweetmemeBot/2.11; +http://tweetmeme.com/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`SearQuBot/SearQuBot v1.0`)
    if err != nil {
        t.Errorf(`Failed to parse 'SearQuBot/SearQuBot v1.0': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Mozilla/5.0 (compatible; ADJUSTbot/2.0; +http://www.ad-just.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; ADJUSTbot/2.0; +http://www.ad-just.jp/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`FTRF: Friendly robot/1.3`)
    if err != nil {
        t.Errorf(`Failed to parse 'FTRF: Friendly robot/1.3': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`kizasi-spider/1.0 (+http://kizasi.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'kizasi-spider/1.0 (+http://kizasi.jp/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`BlogramCrawler/1.0.1(+http://blogram.jp/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'BlogramCrawler/1.0.1(+http://blogram.jp/)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`www2.apserver.net ASP-Ranker Feed Crawler`)
    if err != nil {
        t.Errorf(`Failed to parse 'www2.apserver.net ASP-Ranker Feed Crawler': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Rome Client (http://tinyurl.com/64t5n)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Rome Client (http://tinyurl.com/64t5n)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`UnwindFetchor/1.0 (+http://www.gnip.com/)`)
    if err != nil {
        t.Errorf(`Failed to parse 'UnwindFetchor/1.0 (+http://www.gnip.com/)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)`)
    if err != nil {
        t.Errorf(`Failed to parse 'ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Summify (Summify/1.0.1; +http://summify.com)`)
    if err != nil {
        t.Errorf(`Failed to parse 'Summify (Summify/1.0.1; +http://summify.com)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`PostRank/2.0 (postrank.com; 1 subscribers)`)
    if err != nil {
        t.Errorf(`Failed to parse 'PostRank/2.0 (postrank.com; 1 subscribers)': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`cloudforecastbot`)
    if err != nil {
        t.Errorf(`Failed to parse 'cloudforecastbot': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`SimplePie/1.3-dev (Feed Parser; http://simplepie.org; Allow like Gecko) Build/20111118194710`)
    if err != nil {
        t.Errorf(`Failed to parse 'SimplePie/1.3-dev (Feed Parser; http://simplepie.org; Allow like Gecko) Build/20111118194710': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
    result, err = Parse(`Rainmeter WebParser plugin`)
    if err != nil {
        t.Errorf(`Failed to parse 'Rainmeter WebParser plugin': %s`, err)
    } else {
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
    }
    result, err = Parse(`Data-Hotel-Watchdog/1.1`)
    if err != nil {
        t.Errorf(`Failed to parse 'Data-Hotel-Watchdog/1.1': %s`, err)
    } else {
        if result.Name != "misc crawler" {
            t.Errorf("Expected result.Name to be 'misc crawler', but got '%s'", result.Name)
        }
        if result.Category != "crawler" {
            t.Errorf("Expected result.Category to be 'crawler', but got '%s'", result.Category)
        }
    }
}
