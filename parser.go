package woothee

import (
	"regexp"
	"strings"
)

var (
	rxChromePattern          = regexp.MustCompile(`(?:Chrome|CrMo|CriOS)/([.0-9]+)`)
	rxDocomoVersionPattern   = regexp.MustCompile(`DoCoMo/[.0-9]+[ /]([^- /;()"']+)`)
	rxFirefoxPattern         = regexp.MustCompile(`Firefox/([.0-9]+)`)
	rxFirefoxOSPattern       = regexp.MustCompile(`^Mozilla/[.0-9]+ \((?:Mobile|Tablet);(?:.*;)? rv:([.0-9]+)\) Gecko/[.0-9]+ Firefox/[.0-9]+$`)
	rxFirefoxiOSPattern      = regexp.MustCompile(`FxiOS/([.0-9]+)`)
	rxFOMAVersionPattern     = regexp.MustCompile(`\(([^;)]+);FOMA;`)
	rxGsaPattern             = regexp.MustCompile(`GSA/([.0-9]+)`)
	rxHeadlineReaderPattern  = regexp.MustCompile(`(?i)headline-reader`)
	rxJigPattern             = regexp.MustCompile(`jig browser[^;]+; ([^);]+)`)
	rxKDDIPattern            = regexp.MustCompile(`KDDI-([^- /;()"']+)`)
	rxMaybeRSSPattern        = regexp.MustCompile(`(?i)rss(?:reader|bar|[-_ /;()]|[ +]*/)`)
	rxMaybeCrawlerPattern    = regexp.MustCompile(`(?i)(?:bot|crawler|spider)(?:[-_ ./;@()]|$)`)
	rxMaybeFeedParserPattern = regexp.MustCompile(`(?i)(?:feed|web) ?parser`)
	rxMaybeWatchdogPattern   = regexp.MustCompile(`(?i)watch ?dog`)
	rxMSEdgePattern          = regexp.MustCompile(`(?:Edge|Edg|EdgiOS|EdgA)/([.0-9]+)`)
	rxMSIEPattern            = regexp.MustCompile(`MSIE ([.0-9]+);`)
	rxOperaVersionPattern1   = regexp.MustCompile(`Version/([.0-9]+)`)
	rxOperaVersionPattern2   = regexp.MustCompile(`Opera[/ ]([.0-9]+)`)
	rxOperaVersionPattern3   = regexp.MustCompile(`OPR/([.0-9]+)`)
	rxVivaldiPattern         = regexp.MustCompile(`Vivaldi/([.0-9]+)`)
	rxYandexBrowserPattern   = regexp.MustCompile(`YaBrowser/([.0-9]+)`)
	rxSafariPattern          = regexp.MustCompile(`Version/([.0-9]+)`)
	rxSoftbankPattern        = regexp.MustCompile(`(?:SoftBank|Vodafone|J-PHONE)/[.0-9]+/([^ /;()]+)`)
	rxTridentPattern         = regexp.MustCompile(`Trident/([.0-9]+);(?: BOIE[0-9]+;[A-Z]+;)? rv:([.0-9]+)`)
	rxTridentVersionPattern  = regexp.MustCompile(`rv:([.0-9]+)`)
	rxIEMobilePattern        = regexp.MustCompile(`IEMobile/([.0-9]+);`)
	rxWillcomPattern         = regexp.MustCompile(`(?:WILLCOM|DDIPOCKET);[^/]+/([^ /;()]+)`)
	rxWindowsVersionPattern  = regexp.MustCompile(`Windows ([ .a-zA-Z0-9]+)[;\\)]`)
	rxWinPhone               = regexp.MustCompile(`^Phone(?: OS)? ([.0-9]+)`)
	rxWebviewPattern         = regexp.MustCompile(`iP(hone;|ad;|od) .*like Mac OS X`)
	rxWebviewVersionPattern  = regexp.MustCompile(`Version/([.0-9]+)`)
	rxPPCOsVersion           = regexp.MustCompile(`rv:(\d+\.\d+\.\d+)`)
	rxFreeBSDOsVersion       = regexp.MustCompile(`FreeBSD ([^;\)]+);`)
	rxChromeOSOsVersion      = regexp.MustCompile(`CrOS ([^\)]+)\)`)
	rxAndroidOSOsVersion     = regexp.MustCompile(`Android[- ](\d+(?:\.\d+(?:\.\d+)?)?)`)
	rxPSPOsVersion           = regexp.MustCompile(`PSP \(PlayStation Portable\); ([.0-9]+)\)`)
	rxPS3OsVersion           = regexp.MustCompile(`PLAYSTATION 3;? ([.0-9]+)\)`)
	rxPSVitaOsVersion        = regexp.MustCompile(`PlayStation Vita ([.0-9]+)\)`)
	rxPS4OsVersion           = regexp.MustCompile(`PlayStation 4 ([.0-9]+)\)`)
	rxBlackBerry10OsVersion  = regexp.MustCompile(`BB10(?:.+)Version/([.0-9]+) `)
	rxBlackBerryOsVersion    = regexp.MustCompile(`BlackBerry(?:\d+)/([.0-9]+) `)
)

type Parser struct {
	AgentDataSet DataSet
}

func NewParser() *Parser {
	return &Parser{DefaultDataSet}
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

func (r *Result) populateWith(ds *Result) {
	if ds.Name != "" {
		r.Name = ds.Name
	}

	if ds.Category != "" {
		r.Category = ds.Category
	}

	if ds.Os != "" {
		r.Os = ds.Os
	}

	if ds.Type != "" {
		r.Type = ds.Type
	}

	if ds.Version != "" {
		r.Version = ds.Version
	}

	if ds.Vendor != "" {
		r.Vendor = ds.Vendor
	}
}

func (p *Parser) PopulateDataSet(result *Result, label string) error {
	ds, err := p.LookupDataSet(label)
	if err != nil {
		return err
	}
	result.populateWith(ds)

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

	err = p.ChallengeMsEdge(agent, result)
	if err == nil {
		return
	}

	err = p.ChallengeVivaldi(agent, result)
	if err == nil {
		return
	}

	err = p.ChallengeYandexBrowser(agent, result)
	if err == nil {
		return
	}

	err = p.ChallengeFirefoxiOS(agent, result)
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

	err = p.ChallengeWebview(agent, result)
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

	err = p.ChallengeMiscOs(agent, result)
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

	err = p.ChallengeHTTPLibrary(agent, result)
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
	if !strings.Contains(agent, "Google") {
		return ErrNoMatch
	}

	if strings.Contains(agent, "compatible; Googlebot") {
		if strings.Contains(agent, "compatible; Googlebot-Mobile") {
			return p.PopulateDataSet(result, "GoogleBotMobile")
		}
		return p.PopulateDataSet(result, "GoogleBot")
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
		case strings.Contains(agent, "Y!J-BRZ/YATSHA crawler") || strings.Contains(agent, "Y!J-BRY/YATSH crawler"):
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

	if strings.Contains(agent, "BingPreview") {
		return p.PopulateDataSet(result, "BingPreview")
	}

	if strings.Contains(agent, "Baidu") {
		if strings.Contains(agent, "compatible; Baiduspider") || strings.Contains(agent, "Baiduspider+") || strings.Contains(agent, "Baiduspider-image+") {
			return p.PopulateDataSet(result, "Baiduspider")
		}
	}

	if strings.Contains(agent, "Yeti") {
		if strings.Contains(agent, "http://help.naver.com/robots") || strings.Contains(agent, "http://help.naver.com/support/robots.html") || strings.Contains(agent, "http://naver.me/bot") {
			return p.PopulateDataSet(result, "Yeti")
		}
	}

	if strings.Contains(agent, "FeedBurner/") {
		return p.PopulateDataSet(result, "FeedBurner")
	}

	if strings.Contains(agent, "facebookexternalhit") {
		return p.PopulateDataSet(result, "facebook")
	}

	if strings.Contains(agent, "Twitterbot/") {
		return p.PopulateDataSet(result, "twitter")
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

	if strings.Contains(agent, "trendictionbot") {
		return p.PopulateDataSet(result, "trendictionbot")
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

		result.Os = data.Name
		result.Category = data.Category
		return nil
	}

	return ErrNoMatch
}

func (p *Parser) ChallengeSleipnir(agent string, result *Result) error {
	start := strings.Index(agent, "Sleipnir/")
	if start < 0 {
		return ErrNoMatch
	}

	// Absolutely refuse to use regexps
	agentLen := len(agent)

	end := start + 9
	for ; end < agentLen; end++ {
		switch agent[end] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			// no op
		default:
			break
		}
	}
	version := agent[start+9 : end]
	if version == "" {
		version = ValueUnknown
	}

	err := p.PopulateDataSet(result, "Sleipnir")
	if err != nil {
		return err
	}

	win, err := p.LookupDataSet("Win")
	if err != nil {
		return err
	}

	result.Version = version
	result.Category = win.Category
	result.Os = win.Name

	return nil
}

func (p *Parser) ChallengeHTTPLibrary(agent string, result *Result) error {
	var version string

	if strings.HasPrefix(agent, "Apache-HttpClient/") || strings.HasPrefix(agent, "Jakarta Commons-HttpClient/") || strings.HasPrefix(agent, "Java/") {
		version = "Java"
	} else if i := strings.Index(agent, "HttpClient"); i > -1 {
		// i should be greater than 0 so we can check for either "-" or " "
		// preceding the "HttpClient" string
		if i == 0 {
			return ErrNoMatch
		}

		switch agent[i-1] {
		case '-', ' ':
			// OK, no op
		default:
			return ErrNoMatch
		}

		agentLen := len(agent)
		if agentLen > i+10 { // Longer than "...HttpClient"
			if agentLen > i+11 && agent[i+11] == '/' {
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
	} else if strings.HasPrefix(agent, "curl/") {
		version = "curl"
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

var PHPPrefixPatterns = []string{
	"PHP",
	"WordPress",
	"CakePHP",
	"PukiWiki",
	"PECL::HTTP",
}

var PearPatterns = []string{
	"PEAR HTTP_Request",
	"HTTP_Request",
}

func (p *Parser) isPHP(agent string) bool {
	agentLen := len(agent)
	for _, pattern := range PHPPrefixPatterns {
		if !strings.HasPrefix(agent, pattern) {
			continue
		}

		// Either this pattern is followed by "/", " ", or is exactly same
		// as pattern
		patternLen := len(pattern)
		if patternLen == agentLen {
			return true
		} else if c := agent[patternLen]; c == '/' || c == ' ' {
			return true
		}
	}

	// None of the above patterns matched
	for _, pattern := range PearPatterns {
		i := strings.Index(agent, pattern)
		if i == -1 {
			continue
		}

		// match. this pattern must be followed by either
		// " class" or "2"
		patternLen := len(pattern)
		if i+patternLen+6 <= agentLen {
			if agent[i+patternLen:i+patternLen+6] == " class" {
				return true
			}
		}

		if i+patternLen+1 <= agentLen {
			if agent[i+patternLen] == '2' {
				return true
			}
		}
	}

	return false
}

func (p *Parser) ChallengeMaybeRssReader(agent string, result *Result) error {
	if rxMaybeRSSPattern.MatchString(agent) || rxHeadlineReaderPattern.MatchString(agent) || strings.Contains(agent, "cococ/") {
		return p.PopulateDataSet(result, "VariousRSSReader")
	}

	return ErrNoMatch
}

func (p *Parser) ChallengeMaybeCrawler(agent string, result *Result) error {
	if rxMaybeCrawlerPattern.MatchString(agent) || p.hasCrawlerPrefix(agent) || rxMaybeFeedParserPattern.MatchString(agent) || rxMaybeWatchdogPattern.MatchString(agent) {
		return p.PopulateDataSet(result, "VariousCrawler")
	}
	return ErrNoMatch
}

var CrawlerPrefixPatterns = []string{
	"Rome Client ",
	"UnwindFetchor/",
	"ia_archiver ",
	"Summify ",
	"PostRank/",
}

func (p *Parser) hasCrawlerPrefix(agent string) bool {
	for _, pattern := range CrawlerPrefixPatterns {
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
	var data *Result
	var err error
	var osVersion string

	switch {
	case strings.Contains(agent, "PSP (PlayStation Portable)"):
		data, err = p.LookupDataSet("PSP")
		if matches := rxPSPOsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
	case strings.Contains(agent, "PlayStation Vita"):
		data, err = p.LookupDataSet("PSVita")
		if matches := rxPSVitaOsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
	case strings.Contains(agent, "PLAYSTATION 3 ") || strings.Contains(agent, "PLAYSTATION 3;"):
		data, err = p.LookupDataSet("PS3")
		if matches := rxPS3OsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
	case strings.Contains(agent, "PlayStation 4 "):
		data, err = p.LookupDataSet("PS4")
		if matches := rxPS4OsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
	default:
		return ErrNoMatch
	}

	if err != nil {
		return err
	}

	result.populateWith(data)
	if osVersion != "" {
		result.OsVersion = osVersion
	}

	return nil
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
		}
		return p.PopulateDataSet(result, "Xbox360")
	}

	win, err := p.LookupDataSet("Win")
	if err != nil {
		return err
	}

	match := rxWindowsVersionPattern.FindStringSubmatchIndex(agent)
	if match == nil {
		result.Category = win.Category
		result.Os = win.Name
		return nil
	}

	version := agent[match[2]:match[3]]
	switch version {
	case "NT 10.0":
		win, err = p.LookupDataSet("Win10")
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
		if matches := rxWinPhone.FindStringSubmatch(version); matches != nil {
			version = matches[1]
			win, err = p.LookupDataSet("WinPhone")
		}
	}

	if err != nil {
		return err
	}

	result.Category = win.Category
	result.Os = win.Name
	if version != "" {
		result.OsVersion = version
	}

	return nil
}

var reOSXiPhoneOsVersion = regexp.MustCompile(`; CPU(?: iPhone)? OS (\d+\_\d+(?:_\d+)?) like Mac OS X`)
var reOSXOsVersion = regexp.MustCompile(`Mac OS X (10[._]\d+(?:[._]\d+)?)(?:\)|;)`)

func (p *Parser) ChallengeOsx(agent string, result *Result) error {
	if !strings.Contains(agent, "Mac OS X") {
		return ErrNoMatch
	}

	version := ""

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
		case strings.Contains(agent, "iPod"):
			data, err = p.LookupDataSet("iPod")
		}
		if err != nil {
			return err
		}

		if matches := reOSXiPhoneOsVersion.FindStringSubmatchIndex(agent); matches != nil {
			version = strings.Replace(agent[matches[2]:matches[3]], "_", ".", -1)
		}
	} else {
		if matches := reOSXOsVersion.FindStringSubmatchIndex(agent); matches != nil {
			version = strings.Replace(agent[matches[2]:matches[3]], "_", ".", -1)
		}
	}

	result.Category = data.Category
	result.Os = data.Name
	if version != "" {
		result.OsVersion = version
	}

	return nil
}

func (p *Parser) ChallengeLinux(agent string, result *Result) error {
	if !strings.Contains(agent, "Linux") {
		return ErrNoMatch
	}

	var data *Result
	var err error
	var osVersion string
	if strings.Contains(agent, "Android") {
		data, err = p.LookupDataSet("Android")
		if matches := rxAndroidOSOsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
	} else {
		data, err = p.LookupDataSet("Linux")
	}

	if err != nil {
		return err
	}

	result.Category = data.Category
	result.Os = data.Name
	if osVersion != "" {
		result.OsVersion = osVersion
	}

	return nil
}

func (p *Parser) ChallengeSmartphone(agent string, result *Result) error {
	var data *Result
	var err error
	var osVersion string

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
	case strings.Contains(agent, "BB10"):
		if matches := rxBlackBerry10OsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
		data, err = p.LookupDataSet("BlackBerry10")
		result.Version = ValueUnknown
	case strings.Contains(agent, "BlackBerry"):
		if matches := rxBlackBerryOsVersion.FindStringSubmatch(agent); matches != nil {
			osVersion = matches[1]
		}
		data, err = p.LookupDataSet("BlackBerry")
	}

	if err != nil {
		return err
	}

	firefox, err := p.LookupDataSet("Firefox")
	if err != nil {
		return err
	}

	if result.Name == firefox.Name {
		// Firefox OS specific pattern
		// http://lawrencemandel.com/2012/07/27/decision-made-firefox-os-user-agent-string/
		// https://github.com/woothee/woothee/issues/2
		if matches := rxFirefoxOSPattern.FindStringSubmatch(agent); matches != nil {
			if len(matches) > 1 {
				data, err = p.LookupDataSet("FirefoxOS")
				osVersion = matches[1]
			}
		}
	}

	if data == nil {
		return ErrNoMatch
	}

	result.Category = data.Category
	result.Os = data.Name
	if osVersion != "" {
		result.OsVersion = osVersion
	}

	return nil
}

func (p *Parser) ChallengeMobilephone(agent string, result *Result) error {
	if strings.Contains(agent, "KDDI-") {
		if match := rxKDDIPattern.FindStringSubmatchIndex(agent); match != nil {
			term := agent[match[2]:match[3]]
			data, err := p.LookupDataSet("au")
			if err != nil {
				return err
			}
			result.Category = data.Category
			result.Os = data.Os
			result.Version = term
			return nil
		}
	}

	if strings.Contains(agent, "WILLCOM") || strings.Contains(agent, "DDIPOCKET") {
		if match := rxWillcomPattern.FindStringSubmatchIndex(agent); match != nil {
			term := agent[match[2]:match[3]]
			data, err := p.LookupDataSet("willcom")
			if err != nil {
				return err
			}
			result.Category = data.Category
			result.Os = data.Os
			result.Version = term
			return nil
		}
	}

	if strings.Contains(agent, "SymbianOS") {
		data, err := p.LookupDataSet("SymbianOS")
		if err != nil {
			return err
		}
		result.Category = data.Category
		result.Os = data.Os
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
		result.Os = data.Os
		return nil
	}

	if strings.Contains(agent, "Nintendo Wii;") {
		data, err := p.LookupDataSet("NintendoWii")
		if err != nil {
			return err
		}
		result.Category = data.Category
		result.Os = data.Os
		return nil
	}

	return ErrNoMatch
}

func (p *Parser) ChallengeMsie(agent string, result *Result) error {
	if !strings.Contains(agent, "compatible; MSIE") && !strings.Contains(agent, "Trident/") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if matches := rxMSIEPattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	} else if matches := rxTridentPattern.FindStringSubmatch(agent); matches != nil {
		version = matches[2]
	} else if matches := rxTridentVersionPattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	} else if matches := rxIEMobilePattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	}

	err := p.PopulateDataSet(result, "MSIE")
	if err != nil {
		return err
	}

	result.Version = version
	return nil
}

func (p *Parser) ChallengeMsEdge(agent string, result *Result) error {
	version := ValueUnknown
	if matches := rxMSEdgePattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	} else {
		return ErrNoMatch
	}

	err := p.PopulateDataSet(result, "Edge")
	if err != nil {
		return err
	}
	result.Version = version
	return nil
}

func (p *Parser) ChallengeVivaldi(agent string, result *Result) error {
	version := ValueUnknown
	if matches := rxVivaldiPattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	} else {
		return ErrNoMatch
	}

	err := p.PopulateDataSet(result, "Vivaldi")
	if err != nil {
		return err
	}
	result.Version = version
	return nil
}

func (p *Parser) ChallengeYandexBrowser(agent string, result *Result) error {
	version := ValueUnknown
	if matches := rxYandexBrowserPattern.FindStringSubmatch(agent); matches != nil {
		version = matches[1]
	} else {
		return ErrNoMatch
	}

	err := p.PopulateDataSet(result, "YaBrowser")
	if err != nil {
		return err
	}
	result.Version = version
	return nil
}

func (p *Parser) ChallengeFirefoxiOS(agent string, result *Result) error {
	if matches := rxFirefoxiOSPattern.FindStringSubmatch(agent); matches != nil {
		result.Version = matches[1]
	} else {
		return ErrNoMatch
	}

	err := p.PopulateDataSet(result, "Firefox")
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) ChallengeSafariChrome(agent string, result *Result) error {
	if !strings.Contains(agent, "Safari/") {
		return ErrNoMatch
	}
	if strings.Contains(agent, "Chrome") && strings.Contains(agent, "wv") {
		return ErrNoMatch
	}

	if match := rxChromePattern.FindStringSubmatchIndex(agent); match != nil {
		// Work with Opera (blink)
		if operaMatch := rxOperaVersionPattern3.FindStringSubmatchIndex(agent); operaMatch != nil {
			version := agent[operaMatch[2]:operaMatch[3]]
			err := p.PopulateDataSet(result, "Opera")
			if err != nil {
				return err
			}
			result.Version = version
			return nil
		}

		err := p.PopulateDataSet(result, "Chrome")
		if err != nil {
			return err
		}
		version := agent[match[2]:match[3]]
		result.Version = version
		return nil
	}

	if match := rxGsaPattern.FindStringSubmatchIndex(agent); match != nil {
		version := agent[match[2]:match[3]]
		err := p.PopulateDataSet(result, "GSA")
		if err != nil {
			return err
		}
		result.Version = version
		return nil
	}

	version := ValueUnknown
	if match := rxSafariPattern.FindStringSubmatchIndex(agent); match != nil {
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
	if !strings.Contains(agent, "Firefox/") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxFirefoxPattern.FindStringSubmatchIndex(agent); match != nil {
		version = agent[match[2]:match[3]]
	}

	err := p.PopulateDataSet(result, "Firefox")
	if err != nil {
		return err
	}
	result.Version = version
	return nil
}

func (p *Parser) ChallengeOpera(agent string, result *Result) error {
	if !strings.Contains(agent, "Opera") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxOperaVersionPattern1.FindStringSubmatchIndex(agent); match != nil {
		version = agent[match[2]:match[3]]
	} else if match := rxOperaVersionPattern2.FindStringSubmatchIndex(agent); match != nil {
		version = agent[match[2]:match[3]]
	}

	err := p.PopulateDataSet(result, "Opera")
	if err != nil {
		return err
	}
	result.Version = version

	return nil
}

func (p *Parser) ChallengeWebview(agent string, result *Result) error {
	if strings.Contains(agent, "Chrome") && strings.Contains(agent, "wv") {
		err := p.PopulateDataSet(result, "Webview")
		if err != nil {
			return err
		}

		if matches := rxWebviewVersionPattern.FindStringSubmatch(agent); matches != nil {
			result.Version = matches[1]
		}

		return nil
	}

	if match := rxWebviewPattern.FindStringSubmatchIndex(agent); match == nil || strings.Contains(agent, "Safari/") {
		return ErrNoMatch
	}

	err := p.PopulateDataSet(result, "Webview")
	if err != nil {
		return err
	}

	if matches := rxWebviewVersionPattern.FindStringSubmatch(agent); matches != nil {
		result.Version = matches[1]
	}

	return nil
}

func (p *Parser) ChallengeDocomo(agent string, result *Result) error {
	if !strings.Contains(agent, "DoCoMo") && !strings.Contains(agent, ";FOMA;") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxDocomoVersionPattern.FindStringSubmatchIndex(agent); match != nil {
		version = agent[match[2]:match[3]]
	} else if match := rxFOMAVersionPattern.FindStringSubmatchIndex(agent); match != nil {
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
	if !strings.Contains(agent, "KDDI-") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxKDDIPattern.FindStringSubmatchIndex(agent); match != nil {
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
	if !strings.Contains(agent, "SoftBank") && !strings.Contains(agent, "Vodafone") && !strings.Contains(agent, "J-PHONE") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxSoftbankPattern.FindStringSubmatchIndex(agent); match != nil {
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
	if !strings.Contains(agent, "WILLCOM") && !strings.Contains(agent, "DDIPOCKET") {
		return ErrNoMatch
	}

	version := ValueUnknown
	if match := rxWillcomPattern.FindStringSubmatchIndex(agent); match != nil {
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

		if match := rxJigPattern.FindStringSubmatchIndex(agent); match != nil {
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

func (p *Parser) ChallengeMiscOs(agent string, result *Result) error {
	var data *Result
	var err error

	switch {
	case strings.Contains(agent, "(Win98;"):
		data, err = p.LookupDataSet("Win98")
		result.OsVersion = "98"
	case strings.Contains(agent, "Macintosh; U; PPC;") || strings.Contains(agent, "Mac_PowerPC"):
		data, err = p.LookupDataSet("MacOS")
		if matches := rxPPCOsVersion.FindStringSubmatch(agent); matches != nil {
			result.OsVersion = matches[1]
		}
	case strings.Contains(agent, "X11; FreeBSD "):
		data, err = p.LookupDataSet("BSD")
		if matches := rxFreeBSDOsVersion.FindStringSubmatch(agent); matches != nil {
			result.OsVersion = matches[1]
		}
	case strings.Contains(agent, "X11; CrOS "):
		data, err = p.LookupDataSet("ChromeOS")
		if matches := rxChromeOSOsVersion.FindStringSubmatch(agent); matches != nil {
			result.OsVersion = matches[1]
		}
	}

	if err != nil {
		return err
	}

	if data == nil {
		return ErrNoMatch
	}

	result.Category = data.Category
	result.Os = data.Name

	return nil
}
