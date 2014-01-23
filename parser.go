package woothee

import (
  "regexp"
  "strings"
)

var (
  CHROME_PATTERN = regexp.MustCompile(`(?:Chrome|CrMo|CriOS)/([.0-9]+)`)
  DOCOMO_VERSION_PATTERN = regexp.MustCompile(`DoCoMo/[.0-9]+[ /]([^- /;()"']+)`)
  FIREFOX_PATTERN = regexp.MustCompile(`Firefox/([.0-9]+)`)
  FIREFOXOS_PATTERN = regexp.MustCompile(`^Mozilla/[.0-9]+ \(Mobile;(.*;)? rv:[.0-9]+\) Gecko/[.0-9]+ Firefox/[.0-9]+$`)
  FOMA_VERSION_PATTERN = regexp.MustCompile(`\(([^;)]+);FOMA;`)
  HEADLINE_READER_PATTERN = regexp.MustCompile(`(?i)headline-reader`)
  JIG_PATTERN = regexp.MustCompile(`jig browser[^;]+; ([^);]+)`)
  KDDI_PATTERN = regexp.MustCompile(`KDDI-([^- /;()"']+)`)
  MAYBE_RSS_PATTERN = regexp.MustCompile(`(?i)rss(?:reader|bar|[-_ /;()]|[ +]*/)`)
  MAYBE_CRAWLER_PATTERN = regexp.MustCompile(`(?i)(?:bot|crawler|spider)(?:[-_ ./;@()]|$)`)
  MAYBE_FEED_PARSER_PATTERN = regexp.MustCompile(`(?i)(?:feed|web) ?parser`)
  MAYBE_WATCHDOG_PATTERN = regexp.MustCompile(`(?i)watch ?dog`)
  MSIE_PATTERN = regexp.MustCompile(`MSIE ([.0-9]+);`)
  SAFARI_PATTERN = regexp.MustCompile(`Version/([.0-9]+)`)
  SOFTBANK_PATTERN = regexp.MustCompile(`(?:SoftBank|Vodafone|J-PHONE)/[.0-9]+/([^ /;()]+)`)
  TRIDENT_PATTERN = regexp.MustCompile(`Trident/([.0-9]+); rv ([.0-9]+)`)
  WILLCOM_PATTERN = regexp.MustCompile(`(?:WILLCOM|DDIPOCKET);[^/]+/([^ /;()]+)`)
  WINDOWS_VERSION_PATTERN = regexp.MustCompile(`Windows ([ .a-zA-Z0-9]+)[;\\)]`)
)

type Parser struct {
  AgentDataSet DataSet
}

func NewParser() (*Parser) {
  return &Parser { DefaultDataSet }
}

func (p *Parser) Parse(agent string) (result *Result, err error) {
  result = EmptyResult.Clone()
  if agent == "" || agent == "-" {
    return
  }

  err = p.TryCrawler(agent, result)
  if err == nil {
    return
  }

  err = p.TryBrowser(agent, result)
  if err == nil {
    p.TryOs(agent, result)
    return
  }

  err = p.TryMobilephone(agent, result)
  if err == nil {
    return
  }

  err = p.TryAppliance(agent, result)
  if err == nil {
    return
  }

  err = p.TryMisc(agent, result)
  if err == nil {
    return
  }

  err = p.TryOs(agent, result)
  if err == nil {
    return
  }

  err = p.TryRareCases(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) PopulateDataSet(result *Result, label string) error {
  ds, err := p.LookupDataSet(label)
  if err != nil {
    return err
  }

  result.Name     = ds.Name
  result.Category = ds.Category
  result.Os       = ds.Os
  result.Type     = ds.Type
  result.Version  = ds.Version
  result.Vendor   = ds.Vendor

  return nil
}

func (p *Parser) LookupDataSet(label string) (*Result, error) {
  result := p.AgentDataSet[label]
  if result == nil {
    return nil, ErrNoDataSet
  }

  return result, nil
}

func (p *Parser) TryCrawler(agent string, result *Result) (err error) {
  err = p.ChallengeGoogle(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeCrawlers(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryBrowser(agent string, result *Result) (err error) {
  err = p.ChallengeMsie(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeSafariChrome(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeFirefox(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeOpera(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryMobilephone(agent string, result *Result) (err error) {
  err = p.ChallengeDocomo(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeAu(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeSoftbank(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeWillcom(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeMiscMobilephone(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryAppliance(agent string, result *Result) (err error) {
  err = p.ChallengePlaystation(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeNintendo(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeDigitalTv(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryMisc(agent string, result *Result) (err error) {
  err = p.ChallengeDesktopTools(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryOs(agent string, result *Result) (err error) {
  err = p.ChallengeWindows(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeOsx(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeLinux(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeSmartphone(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeMobilephone(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeAppliance(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) TryRareCases(agent string, result *Result) (err error) {
  err = p.ChallengeSmartphonePatterns(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeSleipnir(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeHttpLibrary(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeMaybeRssReader(agent, result)
  if err == nil {
    return
  }

  err = p.ChallengeMaybeCrawler(agent, result)
  if err == nil {
    return
  }

  err = ErrNoMatch
  return
}

func (p *Parser) ChallengeGoogle(agent string, result *Result) error {
  if ! strings.Contains(agent, "Google") {
    return ErrNoMatch
  }

  if strings.Contains(agent, "compatible; Googlebot") {
    if strings.Contains(agent, "compatible; Googlebot-Mobile") {
      return p.PopulateDataSet(result, "GoogleBotMobile")
    } else {
      return p.PopulateDataSet(result, "GoogleBot")
    }
  }

  if strings.Contains(agent, "Googlebot-Image/") {
    return p.PopulateDataSet(result, "GoogleBot")
  }

  if strings.Contains(agent, "Mediapartners-Google") {
    if strings.Contains(agent, "compatible; Mediapartners-Google") || agent == "Mediapartners-Google" {
      return p.PopulateDataSet(result, "GoogleMediaPartners")
    }
  }

  if strings.Contains(agent, "Feedfetcher-Google;") {
    return p.PopulateDataSet(result, "GoogleFeedFetcher")
  }

  if strings.Contains(agent, "AppEngine-Google") {
    return p.PopulateDataSet(result, "GoogleAppEngine")
  }

  if strings.Contains(agent, "Google Web Preview") {
    return p.PopulateDataSet(result, "GoogleWebPreview")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeCrawlers(agent string, result *Result) error {

  if strings.Contains(agent, "Yahoo") || strings.Contains(agent, "help.yahoo.co.jp/help/jp/") || strings.Contains(agent, "listing.yahoo.co.jp/support/faq/") {
    switch {
    case strings.Contains(agent, "compatible; Yahoo! Slurp"):
      return p.PopulateDataSet(result, "YahooSlurp")
    case strings.Contains(agent, "YahooFeedSeekerJp") || strings.Contains(agent, "YahooFeedSeekerBetaJp"):
      return p.PopulateDataSet(result, "YahooJP")
    case strings.Contains(agent, "crawler (http://listing.yahoo.co.jp/support/faq/") || strings.Contains(agent, "crawler (http://help.yahoo.co.jp/help/jp/"):
      return p.PopulateDataSet(result, "YahooJP")
    case strings.Contains(agent, "Yahoo Pipes"):
      return p.PopulateDataSet(result, "YahooPipes")
    }
  }

  if strings.Contains(agent, "msnbot") {
    return p.PopulateDataSet(result, "msnbot")
  }

  if strings.Contains(agent, "bingbot") {
    if strings.Contains(agent, "compatible; bingbot") {
      return p.PopulateDataSet(result, "bingbot")
    }
  }

  if strings.Contains(agent, "Baidu") {
    if strings.Contains(agent, "compatible; Baiduspider") || strings.Contains(agent, "Baiduspider+") || strings.Contains(agent, "Baiduspider-image+") {
      return p.PopulateDataSet(result, "Baiduspider")
    }
  }

  if strings.Contains(agent, "Yeti") {
    if strings.Contains(agent, "http://help.naver.com/robots") {
      return p.PopulateDataSet(result, "Yeti")
    }
  }

  if strings.Contains(agent, "FeedBurner/") {
    return p.PopulateDataSet(result, "FeedBurner")
  }

  if strings.Contains(agent, "facebookexternalhit") {
    return p.PopulateDataSet(result, "facebook")
  }

  if strings.Contains(agent, "ichiro") {
    if strings.Contains(agent, "http://help.goo.ne.jp/door/crawler.html") || strings.Contains(agent, "compatible; ichiro/mobile goo;") {
      return p.PopulateDataSet(result, "goo")
    }
  }

  if strings.Contains(agent, "gooblogsearch/") {
    return p.PopulateDataSet(result, "goo")
  }

  if strings.Contains(agent, "Apple-PubSub") {
    return p.PopulateDataSet(result, "ApplePubSub")
  }

  if strings.Contains(agent, "(www.radian6.com/crawler)") {
    return p.PopulateDataSet(result, "radian6")
  }

  if strings.Contains(agent, "Genieo/") {
    return p.PopulateDataSet(result, "Genieo")
  }

  if strings.Contains(agent, "labs.topsy.com/butterfly/") {
    return p.PopulateDataSet(result, "topsyButterfly")
  }

  if strings.Contains(agent, "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot") {
    return p.PopulateDataSet(result, "rogerbot")
  }

  if strings.Contains(agent, "compatible; AhrefsBot/") {
    return p.PopulateDataSet(result, "AhrefsBot")
  }

  if strings.Contains(agent, "livedoor FeedFetcher") || strings.Contains(agent, "Fastladder FeedFetcher") {
    return p.PopulateDataSet(result, "livedoorFeedFetcher")
  }

  if strings.Contains(agent, "Hatena ") {
    if strings.Contains(agent, "Hatena Antenna") || strings.Contains(agent, "Hatena Pagetitle Agent") || strings.Contains(agent, "Hatena Diary RSS") {
      return p.PopulateDataSet(result, "Hatena")
    }
  }

  if strings.Contains(agent, "mixi-check") || strings.Contains(agent, "mixi-crawler") || strings.Contains(agent, "mixi-news-crawler") {
    return p.PopulateDataSet(result, "mixi")
  }

  if strings.Contains(agent, "Indy Library") {
    if strings.Contains(agent, "compatible; Indy Library") {
      return p.PopulateDataSet(result, "IndyLibrary")
    }
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeDesktopTools(agent string, result *Result) error {
  if strings.Contains(agent, "AppleSyndication/") {
    return p.PopulateDataSet(result, "SafariRSSReader")
  }

  if strings.Contains(agent, "compatible; Google Desktop/") {
    return p.PopulateDataSet(result, "GoogleDesktop")
  }

  if strings.Contains(agent, "Windows-RSS-Platform") {
    return p.PopulateDataSet(result, "WindowsRSSReader")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeSmartphonePatterns(agent string, result *Result) error {
  if strings.Contains(agent, "CFNetwork/") {
    // This is like iPhone, but only Category and Name are filled
    data, err := p.LookupDataSet("iOS")
    if err != nil {
      return err
    }

    result.Os       = data.Name
    result.Category = data.Category
    return nil
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeSleipnir (agent string, result *Result) error {
  start := strings.Index(agent, "Sleipnir/");
  if start < 0 {
    return ErrNoMatch
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
  version := agent[start + 9:end]
  if version == "" {
    version = VALUE_UNKNOWN
  }

  err := p.PopulateDataSet(result, "Sleipnir")
  if err != nil {
    return err
  }

  win, err := p.LookupDataSet("Win")
  if err != nil {
    return err
  }

  result.Version  = version
  result.Category = win.Category
  result.Os       = win.Name

  return nil
}

func (p *Parser) ChallengeHttpLibrary (agent string, result *Result) error {
  var version string

  if strings.HasPrefix(agent, "Apache-HttpClient/") || strings.HasPrefix(agent, "Jakarta Commons-HttpClient/") || strings.HasPrefix(agent, "Java/") {
    version = "Java"
  } else if i := strings.Index(agent, "HttpClient"); i > - 1 {
    // i should be greater than 0 so we can check for either "-" or " "
    // preceding the "HttpClient" string
    if i == 0 {
      return ErrNoMatch
    }

    switch agent[i - 1] {
    case '-', ' ':
      // OK, no op
    default:
      return ErrNoMatch
    }

    agent_len := len(agent)
    if agent_len > i + 10 { // Longer than "...HttpClient"
      if agent[i + 11] == '/' {
        goto MATCH_JAVA_MISC
      }
      return ErrNoMatch
    }

MATCH_JAVA_MISC:
    version = "Java"
  } else if strings.Contains(agent, "Java(TM) 2 Runtime Environment,") {
    version = "Java"
  } else if strings.HasPrefix(agent, "Wget/") {
    version = "wget"
  } else if strings.HasPrefix(agent, "libwww-perl") || strings.HasPrefix(agent, "WWW-Mechanize") || strings.HasPrefix(agent, "LWP::Simple") || strings.HasPrefix(agent, "LWP ") || strings.HasPrefix(agent, "lwp-trivial") {
    version = "perl"
  } else if strings.HasPrefix(agent, "Ruby") || strings.HasPrefix(agent, "feedzirra") || strings.HasPrefix(agent, "Typhoeus") {
    version = "ruby"
  } else if strings.HasPrefix(agent, "Python-urllib/") || strings.HasPrefix(agent, "Twisted ") {
    version = "python"
  } else if p.isPHP(agent) {
    version = "php"
  }

  if version == "" {
    return ErrNoMatch
  }

  err := p.PopulateDataSet(result, "HTTPLibrary")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
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
    if i + pattern_len + 6 <= agent_len {
      if agent[i + pattern_len:i + pattern_len + 6] == " class" {
        return true
      }
    }

    if i + pattern_len + 1 <= agent_len {
      if agent[i + pattern_len] == '2' {
        return true
      }
    }
  }

  return false
}

func (p *Parser) ChallengeMaybeRssReader(agent string, result *Result) error {
  if MAYBE_RSS_PATTERN.MatchString(agent) || HEADLINE_READER_PATTERN.MatchString(agent) || strings.Contains(agent, "cococ/") {
    return p.PopulateDataSet(result, "VariousRSSReader")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeMaybeCrawler(agent string, result *Result) error {
  if MAYBE_CRAWLER_PATTERN.MatchString(agent) || p.hasCrawlerPrefix(agent) || MAYBE_FEED_PARSER_PATTERN.MatchString(agent) || MAYBE_WATCHDOG_PATTERN.MatchString(agent) {
    return p.PopulateDataSet(result, "VariousCrawler")
  }
  return ErrNoMatch
}

var CRAWLER_PREFIX_PATTERNS []string = []string{
  "Rome Client ",
  "UnwindFetchor/",
  "ia_archiver ",
  "Summify ",
  "PostRank/",
}
func (p *Parser) hasCrawlerPrefix(agent string) bool {
  for _, pattern := range CRAWLER_PREFIX_PATTERNS {
    if strings.HasPrefix(agent, pattern) {
      return true
    }
  }

  if strings.Contains(agent, "ASP-Ranker Feed Crawler") {
    return true
  }

  return false
}

func (p *Parser) ChallengePlaystation(agent string, result *Result) error {
  if strings.Contains(agent, "PSP (PlayStation Portable)") {
    return p.PopulateDataSet(result, "PSP")
  }

  if strings.Contains(agent, "PlayStation Vita") {
    return p.PopulateDataSet(result, "PSVita")
  }

  if strings.Contains(agent, "PLAYSTATION 3 ") || strings.Contains(agent, "PLAYSTATION 3;") {
    return p.PopulateDataSet(result, "PS3")
  }

  if strings.Contains(agent, "PlayStation 4 ") {
    return p.PopulateDataSet(result, "PS4")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeNintendo(agent string, result *Result) error {
  if strings.Contains(agent, "Nintendo 3DS;") {
    return p.PopulateDataSet(result, "Nintendo3DS")
  }

  if strings.Contains(agent, "Nintendo DSi;") {
    return p.PopulateDataSet(result, "NintendoDSi")
  }

  if strings.Contains(agent, "Nintendo Wii;") {
    return p.PopulateDataSet(result, "NintendoWii")
  }

  if strings.Contains(agent, "(Nintendo WiiU)") {
    return p.PopulateDataSet(result, "NintendoWiiU")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeDigitalTv(agent string, result *Result) error {
  if strings.Contains(agent, "InettvBrowser/") {
    return p.PopulateDataSet(result, "DigitalTV")
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeWindows(agent string, result *Result) error {
  if !strings.Contains(agent, "Windows") {
    return ErrNoMatch
  }

  if strings.Contains(agent, "Xbox") {
    if strings.Contains(agent, "Xbox; Xbox One)") {
      return p.PopulateDataSet(result, "XboxOne")
    } else {
      return p.PopulateDataSet(result, "Xbox360")
    }
  }

  win, err := p.LookupDataSet("Win")
  if err != nil {
    return err
  }

  match := WINDOWS_VERSION_PATTERN.FindStringSubmatchIndex(agent)
  if match == nil {
    result.Category = win.Category
    result.Os       = win.Name
    return nil
  }

  version := agent[match[2]:match[3]]
  switch version {
  case "NT 6.3":
    win, err = p.LookupDataSet("Win8.1")
  case "NT 6.2":
    win, err = p.LookupDataSet("Win8")
  case "NT 6.1":
    win, err = p.LookupDataSet("Win7")
  case "NT 6.0":
    win, err = p.LookupDataSet("WinVista")
  case "NT 5.1":
    win, err = p.LookupDataSet("WinXP")
  case "NT 5.0":
    win, err = p.LookupDataSet("Win2000")
  case "NT 4.0":
    win, err = p.LookupDataSet("WinNT4")
  case "98":
    win, err = p.LookupDataSet("Win98")
  case "95":
    win, err = p.LookupDataSet("Win95")
  case "CE":
    win, err = p.LookupDataSet("WinCE")
  default:
    if strings.HasPrefix(version, "Phone OS") {
      win, err = p.LookupDataSet("WinPhone")
    }
  }

  if err != nil {
    return err
  }

  result.Category = win.Category
  result.Os       = win.Name

  return nil
}

func (p *Parser) ChallengeOsx(agent string, result *Result) error {
  if ! strings.Contains(agent, "Mac OS X") {
    return ErrNoMatch
  }

  data, err := p.LookupDataSet("OSX")
  if err != nil {
    return err
  }

  if strings.Contains(agent, "like Mac OS X") {
    switch {
    case strings.Contains(agent, "iPhone;"):
      data, err = p.LookupDataSet("iPhone")
    case strings.Contains(agent, "iPad;"):
      data, err = p.LookupDataSet("iPad")
    case strings.Contains(agent, "iPod;"):
      data, err = p.LookupDataSet("iPod")
    }
    if err != nil {
      return err
    }
  }

  result.Category = data.Category
  result.Os       = data.Name

  return nil
}

func (p *Parser) ChallengeLinux(agent string, result *Result) error {
  if ! strings.Contains(agent, "Linux") {
    return ErrNoMatch
  }

  var data *Result
  var err error
  if strings.Contains(agent, "Android") {
    data, err = p.LookupDataSet("Android")
  } else {
    data, err = p.LookupDataSet("Linux")
  }

  if err != nil {
    return err
  }

  result.Category = data.Category
  result.Os       = data.Name

  return nil
}

func (p *Parser) ChallengeSmartphone(agent string, result *Result) error {
  var data *Result
  var err error

  switch {
  case strings.Contains(agent, "iPhone"):
    data, err = p.LookupDataSet("iPhone")
  case strings.Contains(agent, "iPad"):
    data, err = p.LookupDataSet("iPad")
  case strings.Contains(agent, "iPod"):
    data, err = p.LookupDataSet("iPod")
  case strings.Contains(agent, "Android"):
    data, err = p.LookupDataSet("Android")
  case strings.Contains(agent, "CFNetwork"):
    data, err = p.LookupDataSet("iOS")
  case strings.Contains(agent, "BlackBerry"):
    data, err = p.LookupDataSet("BlackBerry")
  }

  if err != nil {
    return err
  }

  // Firefox OS specific pattern
  // http://lawrencemandel.com/2012/07/27/decision-made-firefox-os-user-agent-string/
  firefox, err := p.LookupDataSet("Firefox")
  if err != nil {
    return err
  }

  if result.Name == firefox.Name {
    if FIREFOXOS_PATTERN.MatchString(agent) {
      data, err = p.LookupDataSet("FirefoxOS")
      if err != nil {
        return err
      }
    }
  }

  if data == nil {
    return ErrNoMatch
  }

  result.Category = data.Category
  result.Os       = data.Name

  return nil
}

func (p *Parser) ChallengeMobilephone(agent string, result *Result) error {
  if strings.Contains(agent, "KDDI-") {
    if match := KDDI_PATTERN.FindStringSubmatchIndex(agent); match != nil {
      term := agent[match[2]:match[3]]
      data, err := p.LookupDataSet("au")
      if err != nil {
        return err
      }
      result.Category = data.Category
      result.Os       = data.Os
      result.Version  = term
      return nil
    }
  }

  if strings.Contains(agent, "WILLCOM") || strings.Contains(agent, "DDIPOCKET") {
    if match := WILLCOM_PATTERN.FindStringSubmatchIndex(agent); match != nil {
      term := agent[match[2]:match[3]]
      data, err := p.LookupDataSet("willcom")
      if err != nil {
        return err
      }
      result.Category = data.Category
      result.Os       = data.Os
      result.Version  = term
      return nil
    }
  }

  if strings.Contains(agent, "SymbianOS") {
    data, err := p.LookupDataSet("SymbianOS")
    if err != nil {
      return err
    }
    result.Category = data.Category
    result.Os       = data.Os
    return nil
  }

  if strings.Contains(agent, "Google Wireless Transcoder") {
    err := p.PopulateDataSet(result, "MobileTranscoder")
    if err != nil {
      return err
    }
    result.Version = "Google"
  }

  if strings.Contains(agent, "Naver Transcoder") {
    err := p.PopulateDataSet(result, "MobileTranscoder")
    if err != nil {
      return err
    }
    result.Version = "Naver"
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeAppliance(agent string, result *Result) error {
  if strings.Contains(agent, "Nintendo DSi;") {
    data, err := p.LookupDataSet("NintendoDSi")
    if err != nil {
      return err
    }
    result.Category = data.Category
    result.Os       = data.Os
    return nil
  }

  if strings.Contains(agent, "Nintendo Wii;") {
    data, err := p.LookupDataSet("NintendoWii")
    if err != nil {
      return err
    }
    result.Category = data.Category
    result.Os       = data.Os
    return nil
  }

  return ErrNoMatch
}

func (p *Parser) ChallengeMsie(agent string, result *Result) error {
  if ! strings.Contains(agent, "compatible; MSIE") && ! strings.Contains(agent, "Trident/") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := MSIE_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  } else if matches := TRIDENT_PATTERN.FindAllStringSubmatchIndex(agent, -1); matches != nil {
    second := matches[1]
    version = agent[second[2]:second[3]]
  }

  err := p.PopulateDataSet(result, "MSIE")
  if err != nil {
    return err
  }

  result.Version = version
  return nil
}

func (p *Parser) ChallengeSafariChrome(agent string, result *Result) error {
  if ! strings.Contains(agent, "Safari/") {
    return ErrNoMatch
  }

  if match := CHROME_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    err := p.PopulateDataSet(result, "Chrome")
    if err != nil {
      return err
    }
    version := agent[match[2]:match[3]]
    result.Version = version
    return nil
  }

  version := VALUE_UNKNOWN
  if match := SAFARI_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "Safari")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

func (p *Parser) ChallengeFirefox(agent string, result *Result) error {
  if ! strings.Contains(agent, "Firefox/") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := FIREFOX_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "Firefox")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

var OPERA_VERSION_PATTERN1 = regexp.MustCompile(`Version/([.0-9]+)`)
var OPERA_VERSION_PATTERN2 = regexp.MustCompile(`Opera[/ ]([.0-9]+)`)
func (p *Parser) ChallengeOpera(agent string, result *Result) error {
  if ! strings.Contains(agent, "Opera") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := OPERA_VERSION_PATTERN1.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  } else if match := OPERA_VERSION_PATTERN2.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "Opera")
  if err != nil {
    return err
  }
  result.Version = version

  return nil
}

func (p *Parser) ChallengeDocomo(agent string, result *Result) error {
  if ! strings.Contains(agent, "DoCoMo") && ! strings.Contains(agent, ";FOMA;") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := DOCOMO_VERSION_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  } else if FOMA_VERSION_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "docomo")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

func (p *Parser) ChallengeAu(agent string, result *Result) error {
  if ! strings.Contains(agent, "KDDI-") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := KDDI_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "au")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

func (p *Parser) ChallengeSoftbank(agent string, result *Result) error {
  if ! strings.Contains(agent, "SoftBank") && ! strings.Contains(agent, "Vodafone") && ! strings.Contains(agent, "J-PHONE") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := SOFTBANK_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "SoftBank")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

func (p *Parser) ChallengeWillcom(agent string, result *Result) error {
  if ! strings.Contains(agent, "WILLCOM") && ! strings.Contains(agent, "DDIPOCKET") {
    return ErrNoMatch
  }

  version := VALUE_UNKNOWN
  if match := WILLCOM_PATTERN.FindStringSubmatchIndex(agent); match != nil {
    version = agent[match[2]:match[3]]
  }

  err := p.PopulateDataSet(result, "willcom")
  if err != nil {
    return err
  }
  result.Version = version
  return nil
}

func (p *Parser) ChallengeMiscMobilephone(agent string, result *Result) error {
  if strings.Contains(agent, "jig browser") {
    err := p.PopulateDataSet(result, "jig")
    if err != nil {
      return err
    }

    if match := JIG_PATTERN.FindStringSubmatchIndex(agent); match != nil {
      result.Version = agent[match[2]:match[3]]
    }
    return nil
  }

  if strings.Contains(agent, "emobile/") || strings.Contains(agent, "OpenBrowser") || strings.Contains(agent, "Browser/Obigo-Browser") {
    err := p.PopulateDataSet(result, "emobile")
    if err != nil {
      return err
    }
    return nil
  }

  if strings.Contains(agent, "SymbianOS") {
    err := p.PopulateDataSet(result, "SymbianOS")
    if err != nil {
      return err
    }
    return nil
  }

  if strings.Contains(agent, "Hatena-Mobile-Gateway/") {
    err := p.PopulateDataSet(result, "MobileTranscoder")
    if err != nil {
      return err
    }
    result.Version = "Hatena"
    return nil
  }

  if strings.Contains(agent, "livedoor-Mobile-Gateway/") {
    err := p.PopulateDataSet(result, "MobileTranscoder")
    if err != nil {
      return err
    }
    result.Version = "livedoor"
    return nil
  }

  return ErrNoMatch
}
