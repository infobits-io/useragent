package useragent

import (
	"regexp"
	"strings"
)

type UserAgent struct {
	userAgent            string
	deviceType           string
	browser              string
	operatingSystem      string
	device               string
	browserCheck         bool // check if the browser is valid
	operatingSystemCheck bool // check if the operating system is valid
	deviceCheck          bool // check if the device is valid
}

func Parse(userAgent string) *UserAgent {
	// Get the browser
	browser := "unknown"
	browserCheck := true

	for _, browserMap := range browsers {
		browserKey := browserMap[0]
		browserRegExString := browserMap[1]

		browserRegEx := regexp.MustCompile(`(?i)` + browserRegExString)
		if browserRegEx.MatchString(userAgent) {
			browser = browserKey
			if strings.HasPrefix(browserKey, "[Bot]") {
				browserCheck = false
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

	// Check for bot indicators
	operatingSystemCheck := true
	deviceCheck := true

	if operatingSystem == "bot" || operatingSystem == "unknown" {
		operatingSystemCheck = false
	}

	if browser == "unknown" {
		browserCheck = false
	}

	if device == "unknown" {
		deviceCheck = false
	}

	// Get the device type
	deviceType := "desktop"
	if tabletCheckRegEx.MatchString(userAgent) {
		deviceType = "tablet"
	} else if mobileCheckRegEx.MatchString(userAgent) {
		deviceType = "mobile"
	}

	// Return object
	return &UserAgent{
		userAgent:            userAgent,
		deviceType:           deviceType,
		browser:              browser,
		device:               device,
		operatingSystem:      operatingSystem,
		browserCheck:         browserCheck,
		operatingSystemCheck: operatingSystemCheck,
		deviceCheck:          deviceCheck,
	}
}

// UserAgent returns the user agent string.
func (ua *UserAgent) UserAgent() string {
	return ua.userAgent
}

// DeviceType returns the device type of the user agent.
func (ua *UserAgent) DeviceType() string {
	return ua.deviceType
}

// Browser returns the browser of the user agent.
func (ua *UserAgent) Browser() string {
	return ua.browser
}

// Device returns the device of the user agent.
func (ua *UserAgent) Device() string {
	return ua.device
}

// OperatingSystem returns the operating system of the user agent.
func (ua *UserAgent) OperatingSystem() string {
	return ua.operatingSystem
}

// IsBot returns true if the user agent is a bot.
// If includeBrowser is true, the browser is also checked.
func (ua *UserAgent) IsBot(includeBrowser bool) bool {
	return (!ua.browserCheck && includeBrowser) || !ua.operatingSystemCheck || !ua.deviceCheck
}

// IsValid returns true if the user agent is not a bot.
func (ua *UserAgent) IsValid() bool {
	return ua.browserCheck && ua.operatingSystemCheck && ua.deviceCheck
}

// IsBrowserValid returns true if the user agent's browser is valid.
func (ua *UserAgent) IsBrowserValid() bool {
	return ua.browserCheck
}

// IsOperatingSystemValid returns true if the user agent's operating system is valid.
func (ua *UserAgent) IsOperatingSystemValid() bool {
	return ua.operatingSystemCheck
}

// IsDeviceValid returns true if the user agent's device is valid.
func (ua *UserAgent) IsDeviceValid() bool {
	return ua.deviceCheck
}

// IsMobile returns true if the user agent is a mobile device.
func (ua *UserAgent) IsMobile() bool {
	return ua.deviceType == "mobile"
}

// IsTablet returns true if the user agent is a tablet device.
func (ua *UserAgent) IsTablet() bool {
	return ua.deviceType == "tablet"
}

// IsDesktop returns true if the user agent is a desktop device.
func (ua *UserAgent) IsDesktop() bool {
	return ua.deviceType == "desktop"
}

// IsWindows returns true if the user agent is a Windows device.
func (ua *UserAgent) IsWindows() bool {
	return ua.operatingSystem == "windows"
}

// IsLinux returns true if the user agent is a Linux device.
func (ua *UserAgent) IsLinux() bool {
	return ua.operatingSystem == "linux"
}

// IsMacOS returns true if the user agent is a MacOS device.
func (ua *UserAgent) IsMacOS() bool {
	return ua.operatingSystem == "macos"
}

// IsAndroid returns true if the user agent is an Android device.
func (ua *UserAgent) IsAndroid() bool {
	return ua.operatingSystem == "android"
}

// IsIOS returns true if the user agent is an iOS device.
func (ua *UserAgent) IsIOS() bool {
	return ua.operatingSystem == "ios"
}

var (
	mobileCheckRegEx = regexp.MustCompile(
		"(?i)/Mobile|iP(hone|od|ad)|Android|BlackBerry|IEMobile|Kindle|NetFront|Silk-Accelerated|(hpw|web)OS|Fennec|Minimo|Opera M(obi|ini)|Blazer|Dolfin|Dolphin|Skyfire|Zune/",
	)
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
		{"Windows Phone", "Windows Phone"},
		{"Open BSD", "OpenBSD"},
		{"FreeBSD", "FreeBSD"},
		{"NetBSD", "NetBSD"},
		{"Solaris", "Solaris|SunOS"},
		{"Android", "Android"},
		{"Ubuntu", "Ubuntu"},
		{"Suse", "Suse"},
		{"Redhat", "Redhat"},
		{"Fedora", "Fedora"},
		{"Centos", "Centos"},
		{"Linux", "(Linux)|(X11)"},
		{"Mac OS", "(Mac_PowerPC)|(Macintosh)"},
		{"Chrome OS", "CrOS"},
		{"BlackBerry", "BlackBerry"},
		{"QNX", "QNX"},
		{"BeOS", "BeOS"},
		{"OS/2", "OS/2"},
		{"iPhone", "iPhone"},
		{"iPad", "iPad"},
		{"iPod", "iPod"},
		{
			"Search Bot",
			"(nuhk)|(Googlebot)|(Yammybot)|(Openbot)|(Slurp)|(MSNBot)|(Ask Jeeves/Teoma)" +
				"|(ia_archiver)|(Baiduspider)|(FacebookExternalHit)|(Twitterbot)|(Riddler)" +
				"|(LinkedInBot)|(Instagram)|(Pinterest)|(chatgpt)|(openai)|(bingbot)" +
				"|(duckduckbot)|(yandexbot)|(snapchat)|(discordbot)",
		},
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
		"FreeBSD":             "linux",
		"NetBSD":              "linux",
		"Solaris":             "linux",
		"Android":             "android",
		"Ubuntu":              "ubuntu",
		"Suse":                "suse",
		"Redhat":              "redhat",
		"Fedora":              "fedora",
		"Centos":              "centos",
		"Linux":               "linux",
		"Mac OS":              "macos",
		"Chrome OS":           "chromeos",
		"BlackBerry":          "blackberry",
		"QNX":                 "qnx",
		"BeOS":                "beos",
		"OS/2":                "os2",
		"iPhone":              "ios",
		"iPad":                "ios",
		"iPod":                "ios",
		"Search Bot":          "bot",
	}
	browsers = [][]string{
		// Browsers
		{"DuckDuckGo", "ddg"},
		{"Brave", "brave"},
		{"Samsung Internet", "samsungbrowser"},
		{"UC Browser", "ucbrowser"},
		{"Opera Mini", "opera mini"},
		{"Opera Mobile", "opera mobi"},
		{"Yandex", "yabrowser"},
		{"360 Safe", "360ee"},
		{"Vivaldi", "vivaldi"},
		{"Tor Browser", "tor"},
		{"Lynx", "lynx"},
		{"SeaMonkey", "seamonkey"},
		{"Pale Moon", "palemoon"},
		{"Midori", "midori"},
		{"Avast Secure Browser", "avast"},
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
		{"[Bot] OpenAI", "openai"},

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
