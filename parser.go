package woothee

import (
  "strings"
)

type Parser struct {
  AgentDataSet DataSet
}

func NewParser() (*Parser) {
  return &Parser { DefaultDataSet }
}

func (p *Parser) LookupDataset(label string) (*Result, error) {
  result := p.AgentDataSet[label]
  if result == nil {
    return nil, ErrNoDataSet
  }

  return result, nil
}

func (p *Parser) TryCrawler(agent string) (result *Result, err error) {
  result, err = p.ChallengeGoogle(agent)
  if err == nil {
    return
  }

  result, err = p.ChallengeCrawlers(agent)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryMisc(agent string) (result *Result, err error) {
  result, err = p.ChallengeDesktopTools(agent)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) ChallengeGoogle(agent string) (*Result, error) {
  if ! strings.Contains(agent, "Google") {
    return nil, ErrNoMatch
  }

  if strings.Contains(agent, "compatible; Googlebot") {
    if strings.Contains(agent, "compatible; Googlebot-Mobile") {
      return p.LookupDataset("GoogleBotMobile")
    } else {
      return p.LookupDataset("GoogleBot")
    }
  }

  if strings.Contains(agent, "Googlebot-Image/") {
    return p.LookupDataset("GoogleBot")
  }

  if strings.Contains(agent, "Mediapartners-Google") {
    if strings.Contains(agent, "compatible; Mediapartners-Google") || agent == "Mediapartners-Google" {
      return p.LookupDataset("GoogleMediaPartners")
    }
  }

  if strings.Contains(agent, "Feedfetcher-Google;") {
    return p.LookupDataset("GoogleFeedFetcher")
  }

  if strings.Contains(agent, "AppEngine-Google") {
    return p.LookupDataset("GoogleAppEngine")
  }

  if strings.Contains(agent, "Google Web Preview") {
    return p.LookupDataset("GoogleWebPreview")
  }

  return nil, ErrNoMatch
}

func (p *Parser) ChallengeCrawlers(agent string) (*Result, error) {

  if strings.Contains(agent, "Yahoo") || strings.Contains(agent, "help.yahoo.co.jp/help/jp/") || strings.Contains(agent, "listing.yahoo.co.jp/support/faq/") {
    switch {
    case strings.Contains(agent, "compatible; Yahoo! Slurp"):
      return p.LookupDataset("YahooSlurp")
    case strings.Contains(agent, "YahooFeedSeekerJp") || strings.Contains(agent, "YahooFeedSeekerBetaJp"):
      return p.LookupDataset("YahooJP")
    case strings.Contains(agent, "crawler (http://listing.yahoo.co.jp/support/faq/") || strings.Contains(agent, "crawler (http://help.yahoo.co.jp/help/jp/"):
      return p.LookupDataset("YahooJP")
    case strings.Contains(agent, "Yahoo Pipes"):
      return p.LookupDataset("YahooPipes")
    }
  }

  if strings.Contains(agent, "msnbot") {
    return p.LookupDataset("msnbot")
  }

  if strings.Contains(agent, "bingbot") {
    if strings.Contains(agent, "compatible; bingbot") {
      return p.LookupDataset("bingbot")
    }
  }

  if strings.Contains(agent, "Baidu") {
    if strings.Contains(agent, "compatible; Baiduspider") || strings.Contains(agent, "Baiduspider+") || strings.Contains(agent, "Baiduspider-image+") {
      return p.LookupDataset("Baiduspider")
    }
  }

  if strings.Contains(agent, "Yeti") {
    if strings.Contains(agent, "http://help.naver.com/robots") {
      return p.LookupDataset("Yeti")
    }
  }

  if strings.Contains(agent, "FeedBurner/") {
    return p.LookupDataset("FeedBurner")
  }

  if strings.Contains(agent, "facebookexternalhit") {
    return p.LookupDataset("facebook")
  }

  if strings.Contains(agent, "ichiro") {
    if strings.Contains(agent, "http://help.goo.ne.jp/door/crawler.html") || strings.Contains(agent, "compatible; ichiro/mobile goo;") {
      return p.LookupDataset("goo")
    }
  }

  if strings.Contains(agent, "gooblogsearch/") {
    return p.LookupDataset("goo")
  }

  if strings.Contains(agent, "Apple-PubSub") {
    return p.LookupDataset("ApplePubSub")
  }

  if strings.Contains(agent, "(www.radian6.com/crawler)") {
    return p.LookupDataset("radian6")
  }

  if strings.Contains(agent, "Genieo/") {
    return p.LookupDataset("Genieo")
  }

  if strings.Contains(agent, "labs.topsy.com/butterfly/") {
    return p.LookupDataset("topsyButterfly")
  }

  if strings.Contains(agent, "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot") {
    return p.LookupDataset("rogerbot")
  }

  if strings.Contains(agent, "compatible; AhrefsBot/") {
    return p.LookupDataset("AhrefsBot")
  }

  if strings.Contains(agent, "livedoor FeedFetcher") || strings.Contains(agent, "Fastladder FeedFetcher") {
    return p.LookupDataset("livedoorFeedFetcher")
  }

  if strings.Contains(agent, "Hatena ") {
    if strings.Contains(agent, "Hatena Antenna") || strings.Contains(agent, "Hatena Pagetitle Agent") || strings.Contains(agent, "Hatena Diary RSS") {
      return p.LookupDataset("Hatena")
    }
  }

  if strings.Contains(agent, "mixi-check") || strings.Contains(agent, "mixi-crawler") || strings.Contains(agent, "mixi-news-crawler") {
    return p.LookupDataset("mixi")
  }

  if strings.Contains(agent, "Indy Library") {
    if strings.Contains(agent, "compatible; Indy Library") {
      return p.LookupDataset("IndyLibrary")
    }
  }

  return nil, ErrNoMatch
}

func (p *Parser) ChallengeDesktopTools(agent string) (*Result, error) {
  if strings.Contains(agent, "AppleSyndication/") {
    return p.LookupDataset("SafariRSSReader")
  }

  if strings.Contains(agent, "compatible; Google Desktop/") {
    return p.LookupDataset("GoogleDesktop")
  }

  if strings.Contains(agent, "Windows-RSS-Platform") {
    return p.LookupDataset("WindowsRSSReader")
  }

  return nil, ErrNoMatch
}
