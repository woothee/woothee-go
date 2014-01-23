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

func (p *Parser) TryRareCases(agent string) (result *Result, err error) {
  result, err = p.ChallengeSmartphonePatterns(agent)
  if err == nil {
    return
  }

  result, err = p.ChallengeSleipnir(agent)
  if err == nil {
    return
  }

  result, err = p.ChallengeHttpLibrary(agent)
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

func (p *Parser) ChallengeSmartphonePatterns(agent string) (*Result, error) {
  if strings.Contains(agent, "CFNetwork/") {
    // This is like iPhone, but only Category and Name are filled
    result, err := p.LookupDataset("iOS")
    if err != nil {
      return nil, err
    } else {
      r := EmptyResult.Clone()
      r.Name = result.Name
      r.Category = result.Category
      return r, nil
    }
  }

  return nil, ErrNoMatch
}

func (p *Parser) ChallengeSleipnir (agent string) (*Result, error) {
  start := strings.Index(agent, "Sleipnir/");
  if start < 0 {
    return nil, ErrNoMatch
  }

  // Absolutely refuse to use regexps
  agent_len := len(agent)

  end := start + 9
  for ; end < agent_len; end++ {
    switch agent[end] {
    case '0','1','2','3','4','5','6','7','8','9','.':
      // no op
    default:
      break
    }
  }
  version := agent[start:end]
  if version == "" {
    version = VALUE_UNKNOWN
  }

  result, err := p.LookupDataset("Sleipnir")
  if err != nil {
    return nil, err
  }

  win, err := p.LookupDataset("Win")
  if err != nil {
    return nil, err
  }

  r := result.Clone()
  r.Version = version
  r.Category = win.Category
  r.Os = win.Name

  return r, nil
}

func (p *Parser) ChallengeHttpLibrary (agent string) (*Result, error) {
  var version string

  if strings.HasPrefix(agent, "Apache-HttpClient/") || strings.HasPrefix(agent, "Jakarta Commons-HttpClient/") || strings.HasPrefix(agent, "Java/") {
    version = "Java"
  } else if i := strings.Index(agent, "HttpClient"); i > - 1 {
    // i should be greater than 0 so we can check for either "-" or " "
    // preceding the "HttpClient" string
    if i == 0 {
      return nil, ErrNoMatch
    }

    switch agent[i - 1] {
    case '-', ' ':
      // OK, no op
    default:
      return nil, ErrNoMatch
    }

    agent_len := len(agent)
    if agent_len > i + 10 { // Longer than "...HttpClient"
      if agent[i + 11] == '/' {
        goto MATCH_JAVA_MISC
      }
      return nil, ErrNoMatch
    }

MATCH_JAVA_MISC:
    version = "Java"
  } else if strings.Contains(agent, "Java(TM) 2 Runtime Environment,") {
    version = "Java"
  } else if strings.HasPrefix(agent, "Wget/") {
    version = "wget"
  } else if strings.HasPrefix(agent, "libwww-perl") || strings.HasPrefix(agent, "WWW-Mechanize") || strings.HasPrefix(agent, "LWP::Simple") || strings.HasPrefix(agent, "LWP ") || strings.HasPrefix(agent, "lwp-trivial") {
    version = "perl"
  } else if strings.HasPrefix(agent, "Ruby") || strings.HasPrefix(agent, "feedzirra") || strings.HasPrefix(agent, "Typoeus") {
    version = "ruby"
  } else if strings.HasPrefix(agent, "Python-urllib/") || strings.HasPrefix(agent, "Twisted ") {
    version = "python"
  } else if p.isPHP(agent) {
    version = "php"
  }

  if version == "" {
    return nil, ErrNoMatch
  }

  result, err := p.LookupDataset("HTTPLibrary")
  if err != nil {
    return nil, err
  }
  result = result.Clone()
  result.Version = version
  return result, nil
}

var PHP_PREFIX_PATTERNS []string = []string{
  "PHP",
  "WordPress",
  "CakePHP",
  "PukiWiki",
  "PECL::HTTP",
}

var PEAR_PATTERNS []string = []string {
  "PEAR HTTP_Request",
  "HTTP_Request",
}
func (p *Parser) isPHP(agent string) bool {
  agent_len := len(agent)
  for _, pattern := range PHP_PREFIX_PATTERNS {
    if ! strings.HasPrefix(agent, pattern) {
      continue
    }

    // Either this pattern is followed by "/", " ", or is exactly same
    // as pattern
    pattern_len := len(pattern)
    if pattern_len == agent_len {
      return true
    } else if c := agent[pattern_len]; c == '/' || c == ' ' {
      return true
    }
  }

  // None of the above patterns matched
  for _, pattern := range PEAR_PATTERNS {
    i := strings.Index(agent, pattern)
    if i == -1 {
      continue
    }

    // match. this pattern must be followed by either
    // " class" or "2"
    pattern_len := len(pattern)
    if i + pattern_len + 1 <= agent_len {
      if agent[i + pattern_len] == '2' {
        return true
      }
    }

    if i + pattern_len + 6 <= agent_len {
      if agent[i + pattern_len:i + pattern_len + 6] == " class" {
        return true
      }
    }
  }

  return false
}
