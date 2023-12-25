package useragent

import (
	"regexp"
	"strings"
)

// TODO: add some client side checks for duckduckgo and brave
type UserAgent struct {
	userAgent       string
	deviceType      string
	browser         string
	operatingSystem string
	device          string
	bot             bool
}

func Parse(userAgent string) *UserAgent {
	// Get the browser
	browser := "unknown"
	bot := false
	for _, browserMap := range browsers {
		browserKey := browserMap[0]
		browserRegExString := browserMap[1]
		browserRegEx := regexp.MustCompile(`(?i)` + browserRegExString)
		if browserRegEx.MatchString(userAgent) {
			browser = browserKey
			if strings.HasPrefix(browserKey, "[Bot]") {
				bot = true
			}
			break
		}
	}

	// Get the device
	device := "unknown"
	operatingSystem := "unknown"
	for _, deviceMap := range devices {
		deviceKey := deviceMap[0]
		deviceRegExString := deviceMap[1]
		deviceRegEx := regexp.MustCompile(`(?i)` + deviceRegExString)
		if deviceRegEx.MatchString(userAgent) {
			device = deviceKey
			operatingSystem = deviceOperatingSystem[deviceKey]
			break
		}
	}

	// Get the device type
	deviceType := "desktop"
	if tabletCheckRegEx.MatchString(userAgent) {
		deviceType = "tablet"
	} else if mobileCheckRegEx.MatchString(userAgent) {
		deviceType = "mobile"
	}

	// return object
	return &UserAgent{
		userAgent:       userAgent,
		deviceType:      deviceType,
		browser:         browser,
		device:          device,
		operatingSystem: operatingSystem,
		bot:             bot,
	}
}

func (ua *UserAgent) UserAgent() string {
	return ua.userAgent
}

func (ua *UserAgent) DeviceType() string {
	return ua.deviceType
}

func (ua *UserAgent) Browser() string {
	return ua.browser
}

func (ua *UserAgent) Device() string {
	return ua.device
}

func (ua *UserAgent) OperatingSystem() string {
	return ua.operatingSystem
}

func (ua *UserAgent) IsBot() bool {
	return ua.bot
}

var (
	mobileCheckRegEx = regexp.MustCompile("(?i)/Mobile|iP(hone|od|ad)|Android|BlackBerry|IEMobile|Kindle|NetFront|Silk-Accelerated|(hpw|web)OS|Fennec|Minimo|Opera M(obi|ini)|Blazer|Dolfin|Dolphin|Skyfire|Zune/")
	tabletCheckRegEx = regexp.MustCompile("(?i)(tablet|ipad|playbook)|.*mobile.*android.*")
	devices          = [][]string{
		{"Windows 3.11", "Win16"},
		{"Windows 95", "(Windows 95)|(Win95)|(Windows_95)"},
		{"Windows 98", "(Windows 98)|(Win98)"},
		{"Windows 2000", "(Windows NT 5.0)|(Windows 2000)"},
		{"Windows XP", "(Windows NT 5.1)|(Windows XP)"},
		{"Windows Server 2003", "(Windows NT 5.2)"},
		{"Windows Vista", "(Windows NT 6.0)"},
		{"Windows 7", "(Windows NT 6.1)"},
		{"Windows 8", "(Windows NT 6.2)|(WOW64)"},
		{"Windows 10", "(Windows 10.0)|(Windows NT 10.0)"},
		{"Windows NT 4.0", "(Windows NT 4.0)|(WinNT4.0)|(WinNT)|(Windows NT)"},
		{"Windows ME", "Windows ME"},
		{"Open BSD", "OpenBSD"},
		{"Sun OS", "SunOS"},
		{"Android", "Android"},
		{"Ubuntu", "Ubuntu"},
		{"Suse", "Suse"},
		{"Redhat", "Redhat"},
		{"Fedora", "Fedora"},
		{"Centos", "Centos"},
		{"Linux", "(Linux)|(X11)"},
		{"Mac OS", "(Mac_PowerPC)|(Macintosh)"},
		{"QNX", "QNX"},
		{"BeOS", "BeOS"},
		{"OS/2", "OS/2"},
		{"iPhone", "iPhone"},
		{"iPad", "iPad"},
		{"iPod", "iPod"},
		{"Search Bot", "(nuhk)|(Googlebot)|(Yammybot)|(Openbot)|(Slurp)|(MSNBot)|(Ask Jeeves/Teoma)|(ia_archiver)|(Baiduspider)|(FacebookExternalHit)|(Twitterbot)|(Riddler)|(LinkedInBot)|(Instagram)|(Pinterest)|(chatgpt)|(bingbot)|(duckduckbot)|(yandexbot)|(snapchat)|(discordbot)"},
	}
	deviceOperatingSystem = map[string]string{
		"Windows 3.11":        "windows",
		"Windows 95":          "windows",
		"Windows 98":          "windows",
		"Windows 2000":        "windows",
		"Windows XP":          "windows",
		"Windows Server 2003": "windows",
		"Windows Vista":       "windows",
		"Windows 7":           "windows",
		"Windows 8":           "windows",
		"Windows 10":          "windows",
		"Windows NT 4.0":      "windows",
		"Windows ME":          "windows",
		"Open BSD":            "linux",
		"Sun OS":              "linux",
		"Android":             "android",
		"Ubuntu":              "ubuntu",
		"Suse":                "suse",
		"Redhat":              "redhat",
		"Fedora":              "fedora",
		"Centos":              "centos",
		"Linux":               "linux",
		"Mac OS":              "macos",
		"QNX":                 "qnx",
		"BeOS":                "beos",
		"OS/2":                "os2",
		"iPhone":              "ios",
		"iPad":                "ios",
		"iPod":                "ios",
		"Search Bot":          "bot",
	}
	browsers = [][]string{
		{"Opera", "(opera)|(opr/)"},
		{"Edge", "(edge)|(edg)"},
		{"Chrome", "(chrome)|(crios)"},
		{"Safari", "safari"},
		{"Firefox", "firefox"},
		{"Internet Explorer", "(msie)|(trident/7)"},
		// Search Engines
		{"[Bot] Googlebot", "google"},
		{"[Bot] Bingbot", "bing"},
		{"[Bot] Yahoo! Slurp", "slurp"},
		{"[Bot] DuckDuckBot", "(duckduckgo)|(duckduckbot)"},
		{"[Bot] Baidu", "baidu"},
		{"[Bot] Yandex", "yandex"},
		{"[Bot] Sogou", "sogou"},
		{"[Bot] Exabot", "exabot"},
		{"[Bot] MSN", "msn"},
		// Chat bots
		{"[Bot] ChatGPT", "chatgpt"},
		// Social Media
		{"[Bot] Facebook", "facebook"},
		{"[Bot] Pinterest", "pinterest"},
		{"[Bot] LinkedInBot", "linkedin"},
		{"[Bot] Instagram", "instagram"},
		{"[Bot] Twitterbot", "twitter"},
		{"[Bot] Snapchat", "snapchat"},
		{"[Bot] Discord", "discord"},
		// Common Tools and Bots
		{"[Bot] Majestic", "mj12bot"},
		{"[Bot] Ahrefs", "ahrefs"},
		{"[Bot] SEMRush", "semrush"},
		{"[Bot] Moz or OpenSiteExplorer", "(rogerbot)|(dotbot)"},
		{"[Bot] Screaming Frog", "(frog)|(screaming)"},
		{"[Bot] Pingdom", "pingdom"},
		{"[Bot] Riddler", "riddler"},
		{"[Bot] W3C Validator", "w3c_validator"},
		// Miscellaneous

		// Check for strings commonly used in bot user agents
		{"[Bot] Other", "(crawler)|(api)|(spider)|(http)|(bot)|(archive)|(info)|(data)"},
	}
)
