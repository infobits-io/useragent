package useragent

import (
	"regexp"
)

// UserAgent represents a parsed user agent string.
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

// browserPattern holds a pre-compiled regex for matching a browser or bot.
type browserPattern struct {
	name  string
	regex *regexp.Regexp
	isBot bool
}

// devicePattern holds a pre-compiled regex for matching a device/OS.
type devicePattern struct {
	name  string
	regex *regexp.Regexp
	os    string
}

// Parse parses a user agent string and returns a UserAgent.
func Parse(userAgent string) *UserAgent {
	// Get the browser
	browser := "unknown"
	browserCheck := true

	for i := range browsers {
		bp := &browsers[i]
		if bp.regex.MatchString(userAgent) {
			browser = bp.name
			if bp.isBot {
				browserCheck = false
			}

			break
		}
	}

	// Get the device
	device := "unknown"
	operatingSystem := "unknown"

	for i := range devices {
		dp := &devices[i]
		if dp.regex.MatchString(userAgent) {
			device = dp.name
			operatingSystem = dp.os

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

func compileBrowser(name, pattern string, isBot bool) browserPattern {
	return browserPattern{
		name:  name,
		regex: regexp.MustCompile(`(?i)` + pattern),
		isBot: isBot,
	}
}

func compileDevice(name, pattern, os string) devicePattern {
	return devicePattern{
		name:  name,
		regex: regexp.MustCompile(`(?i)` + pattern),
		os:    os,
	}
}

var (
	mobileCheckRegEx = regexp.MustCompile(
		`(?i)Mobile|iP(hone|od|ad)|Android|BlackBerry|IEMobile|Kindle|NetFront|` +
			`Silk-Accelerated|(hpw|web)OS|Fennec|Minimo|Opera M(obi|ini)|Blazer|Dolfin|Dolphin|Skyfire|Zune`,
	)
	tabletCheckRegEx = regexp.MustCompile(`(?i)(tablet|ipad|playbook)|.*mobile.*android.*`)

	devices = [...]devicePattern{
		compileDevice("Windows 3.11", `Win16`, "windows"),
		compileDevice("Windows 95", `(Windows 95)|(Win95)|(Windows_95)`, "windows"),
		compileDevice("Windows 98", `(Windows 98)|(Win98)`, "windows"),
		compileDevice("Windows 2000", `(Windows NT 5.0)|(Windows 2000)`, "windows"),
		compileDevice("Windows XP", `(Windows NT 5.1)|(Windows XP)`, "windows"),
		compileDevice("Windows Server 2003", `(Windows NT 5.2)`, "windows"),
		compileDevice("Windows Vista", `(Windows NT 6.0)`, "windows"),
		compileDevice("Windows 7", `(Windows NT 6.1)`, "windows"),
		compileDevice("Windows 8", `(Windows NT 6.2)`, "windows"),
		compileDevice("Windows 10", `(Windows 10.0)|(Windows NT 10.0)`, "windows"),
		compileDevice("Windows NT 4.0", `(Windows NT 4.0)|(WinNT4.0)|(WinNT)|(Windows NT)`, "windows"),
		compileDevice("Windows ME", `Windows ME`, "windows"),
		compileDevice("Windows Phone", `Windows Phone`, "windows"),
		compileDevice("Open BSD", `OpenBSD`, "linux"),
		compileDevice("FreeBSD", `FreeBSD`, "linux"),
		compileDevice("NetBSD", `NetBSD`, "linux"),
		compileDevice("Solaris", `Solaris|SunOS`, "linux"),
		compileDevice("Android", `Android`, "android"),
		compileDevice("Ubuntu", `Ubuntu`, "ubuntu"),
		compileDevice("Suse", `Suse`, "suse"),
		compileDevice("Redhat", `Redhat`, "redhat"),
		compileDevice("Fedora", `Fedora`, "fedora"),
		compileDevice("Centos", `Centos`, "centos"),
		compileDevice("Chrome OS", `CrOS`, "chromeos"),
		compileDevice("Linux", `(Linux)|(X11)`, "linux"),
		compileDevice("Mac OS", `(Mac_PowerPC)|(Macintosh)`, "macos"),
		compileDevice("BlackBerry", `BlackBerry`, "blackberry"),
		compileDevice("QNX", `QNX`, "qnx"),
		compileDevice("BeOS", `BeOS`, "beos"),
		compileDevice("OS/2", `OS/2`, "os2"),
		compileDevice("iPhone", `iPhone`, "ios"),
		compileDevice("iPad", `iPad`, "ios"),
		compileDevice("iPod", `iPod`, "ios"),
		compileDevice("Search Bot",
			`(nuhk)|(Googlebot)|(Yammybot)|(Openbot)|(Slurp)|(MSNBot)|(Ask Jeeves/Teoma)`+
				`|(ia_archiver)|(Baiduspider)|(FacebookExternalHit)|(Twitterbot)|(Riddler)`+
				`|(LinkedInBot)|(Instagram)|(Pinterest)|(chatgpt)|(openai)|(bingbot)`+
				`|(duckduckbot)|(yandexbot)|(snapchat)|(discordbot)`+
				`|(claudebot)|(gptbot)|(perplexitybot)|(bytespider)|(petalbot)|(applebot)|(amazonbot)`,
			"bot"),
	}

	browsers = [...]browserPattern{
		// Browsers
		compileBrowser("DuckDuckGo", `ddg`, false),
		compileBrowser("Brave", `brave`, false),
		compileBrowser("Samsung Internet", `samsungbrowser`, false),
		compileBrowser("UC Browser", `ucbrowser`, false),
		compileBrowser("Opera Mini", `opera mini`, false),
		compileBrowser("Opera Mobile", `opera mobi`, false),
		compileBrowser("Yandex", `yabrowser`, false),
		compileBrowser("360 Safe", `360ee`, false),
		compileBrowser("Vivaldi", `vivaldi`, false),
		compileBrowser("Arc", `arc/`, false),
		compileBrowser("Opera GX", `oprgx`, false),
		compileBrowser("Tor Browser", `tor`, false),
		compileBrowser("Lynx", `lynx`, false),
		compileBrowser("SeaMonkey", `seamonkey`, false),
		compileBrowser("Pale Moon", `palemoon`, false),
		compileBrowser("Midori", `midori`, false),
		compileBrowser("Avast Secure Browser", `avast`, false),
		compileBrowser("Opera", `(opera)|(opr/)`, false),
		compileBrowser("Edge", `(edge)|(edg)`, false),
		compileBrowser("Chrome", `(chrome)|(crios)`, false),
		compileBrowser("Safari", `safari`, false),
		compileBrowser("Firefox", `firefox`, false),
		compileBrowser("Internet Explorer", `(msie)|(trident/7)`, false),
		// Search Engines
		compileBrowser("[Bot] Googlebot", `google`, true),
		compileBrowser("[Bot] Bingbot", `bing`, true),
		compileBrowser("[Bot] Yahoo! Slurp", `slurp`, true),
		compileBrowser("[Bot] DuckDuckBot", `(duckduckgo)|(duckduckbot)`, true),
		compileBrowser("[Bot] Baidu", `baidu`, true),
		compileBrowser("[Bot] Yandex", `yandex`, true),
		compileBrowser("[Bot] Sogou", `sogou`, true),
		compileBrowser("[Bot] Exabot", `exabot`, true),
		compileBrowser("[Bot] MSN", `msn`, true),
		// Chat bots
		compileBrowser("[Bot] ChatGPT", `chatgpt`, true),
		compileBrowser("[Bot] ClaudeBot", `claudebot`, true),
		compileBrowser("[Bot] GPTBot", `gptbot`, true),
		compileBrowser("[Bot] PerplexityBot", `perplexitybot`, true),
		compileBrowser("[Bot] OpenAI", `openai`, true),
		// Social Media
		compileBrowser("[Bot] Facebook", `facebook`, true),
		compileBrowser("[Bot] Pinterest", `pinterest`, true),
		compileBrowser("[Bot] LinkedInBot", `linkedin`, true),
		compileBrowser("[Bot] Instagram", `instagram`, true),
		compileBrowser("[Bot] Twitterbot", `twitter`, true),
		compileBrowser("[Bot] Snapchat", `snapchat`, true),
		compileBrowser("[Bot] Discord", `discord`, true),
		// Common Tools and Bots
		compileBrowser("[Bot] Bytespider", `bytespider`, true),
		compileBrowser("[Bot] PetalBot", `petalbot`, true),
		compileBrowser("[Bot] Applebot", `applebot`, true),
		compileBrowser("[Bot] Amazon", `amazonbot`, true),
		compileBrowser("[Bot] Majestic", `mj12bot`, true),
		compileBrowser("[Bot] Ahrefs", `ahrefs`, true),
		compileBrowser("[Bot] SEMRush", `semrush`, true),
		compileBrowser("[Bot] Moz or OpenSiteExplorer", `(rogerbot)|(dotbot)`, true),
		compileBrowser("[Bot] Screaming Frog", `(frog)|(screaming)`, true),
		compileBrowser("[Bot] Pingdom", `pingdom`, true),
		compileBrowser("[Bot] Riddler", `riddler`, true),
		compileBrowser("[Bot] W3C Validator", `w3c_validator`, true),
		// Check for strings commonly used in bot user agents
		compileBrowser("[Bot] Other", `(crawler)|(api)|(spider)|(http)|(bot)|(archive)|(info)|(data)`, true),
	}
)
