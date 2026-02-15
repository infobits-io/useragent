package useragent

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name       string
		userAgent  string
		deviceType string
		browser    string
		device     string
		os         string
		isBot      bool
	}{
		// Desktop browsers
		{
			name:       "Chrome on Windows 10",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Firefox on Windows 10",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
			deviceType: "desktop",
			browser:    "Firefox",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Edge on Windows 10",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
			deviceType: "desktop",
			browser:    "Edge",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Safari on macOS",
			userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15",
			deviceType: "desktop",
			browser:    "Safari",
			device:     "Mac OS",
			os:         "macos",
			isBot:      false,
		},
		{
			name:       "Opera on Windows",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/106.0.0.0",
			deviceType: "desktop",
			browser:    "Opera",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Opera GX",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPRGX/106.0.0.0",
			deviceType: "desktop",
			browser:    "Opera GX",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Brave on macOS",
			userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Brave Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Brave",
			device:     "Mac OS",
			os:         "macos",
			isBot:      false,
		},
		{
			name:       "Vivaldi on Linux",
			userAgent:  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Vivaldi/6.5",
			deviceType: "desktop",
			browser:    "Vivaldi",
			device:     "Linux",
			os:         "linux",
			isBot:      false,
		},
		{
			name:       "Arc browser",
			userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Arc/1.22.0 Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Arc",
			device:     "Mac OS",
			os:         "macos",
			isBot:      false,
		},
		{
			name:       "DuckDuckGo browser",
			userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15 DDG/7",
			deviceType: "desktop",
			browser:    "DuckDuckGo",
			device:     "Mac OS",
			os:         "macos",
			isBot:      false,
		},
		{
			name:       "Tor Browser on Linux",
			userAgent:  "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Tor Browser Firefox/115.0",
			deviceType: "desktop",
			browser:    "Tor Browser",
			device:     "Linux",
			os:         "linux",
			isBot:      false,
		},
		{
			name:       "Internet Explorer 11",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
			deviceType: "desktop",
			browser:    "Internet Explorer",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
		// Mobile browsers
		{
			name:       "Chrome on iPhone",
			userAgent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/120.0.6099.119 Mobile/15E148 Safari/604.1",
			deviceType: "mobile",
			browser:    "Chrome",
			device:     "iPhone",
			os:         "ios",
			isBot:      false,
		},
		{
			name:       "Safari on iPhone",
			userAgent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
			deviceType: "mobile",
			browser:    "Safari",
			device:     "iPhone",
			os:         "ios",
			isBot:      false,
		},
		{
			name:       "Chrome on Android phone",
			userAgent:  "Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.144 Mobile Safari/537.36",
			deviceType: "mobile",
			browser:    "Chrome",
			device:     "Android",
			os:         "android",
			isBot:      false,
		},
		{
			name:       "Samsung Internet",
			userAgent:  "Mozilla/5.0 (Linux; Android 13; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/23.0 Chrome/115.0.0.0 Mobile Safari/537.36",
			deviceType: "mobile",
			browser:    "Samsung Internet",
			device:     "Android",
			os:         "android",
			isBot:      false,
		},
		// Tablets
		{
			name:       "iPad with Safari",
			userAgent:  "Mozilla/5.0 (iPad; CPU OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
			deviceType: "tablet",
			browser:    "Safari",
			device:     "iPad",
			os:         "ios",
			isBot:      false,
		},
		{
			name:       "Android tablet",
			userAgent:  "Mozilla/5.0 (Linux; Android 13; SM-X710) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.144 Safari/537.36 Tablet",
			deviceType: "tablet",
			browser:    "Chrome",
			device:     "Android",
			os:         "android",
			isBot:      false,
		},
		// Windows versions
		{
			name:       "Windows 7",
			userAgent:  "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Windows 7",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Windows 8",
			userAgent:  "Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Windows 8",
			os:         "windows",
			isBot:      false,
		},
		{
			name:       "Chrome OS",
			userAgent:  "Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Chrome OS",
			os:         "chromeos",
			isBot:      false,
		},
		{
			name:       "Linux desktop",
			userAgent:  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Linux",
			os:         "linux",
			isBot:      false,
		},
		// Search engine bots
		{
			name:       "Googlebot",
			userAgent:  "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			deviceType: "desktop",
			browser:    "[Bot] Googlebot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		{
			name:       "Bingbot",
			userAgent:  "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			deviceType: "desktop",
			browser:    "[Bot] Bingbot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		// AI bots
		{
			name:       "ClaudeBot",
			userAgent:  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ClaudeBot/1.0; +claudebot@anthropic.com)",
			deviceType: "desktop",
			browser:    "[Bot] ClaudeBot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		{
			name:       "GPTBot",
			userAgent:  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; GPTBot/1.0; +https://openai.com/gptbot)",
			deviceType: "desktop",
			browser:    "[Bot] GPTBot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		{
			name:       "ChatGPT",
			userAgent:  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ChatGPT-User/1.0; +https://openai.com/bot)",
			deviceType: "desktop",
			browser:    "[Bot] ChatGPT",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		{
			name:       "PerplexityBot",
			userAgent:  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; PerplexityBot/1.0; +https://perplexity.ai/perplexitybot)",
			deviceType: "desktop",
			browser:    "[Bot] PerplexityBot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		// Social media bots
		{
			name:       "Facebook bot",
			userAgent:  "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
			deviceType: "desktop",
			browser:    "[Bot] Facebook",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		{
			name:       "LinkedInBot",
			userAgent:  "LinkedInBot/1.0 (compatible; Mozilla/5.0; Apache-HttpClient +http://www.linkedin.com)",
			deviceType: "desktop",
			browser:    "[Bot] LinkedInBot",
			device:     "Search Bot",
			os:         "bot",
			isBot:      true,
		},
		// SEO bots
		{
			name:       "Ahrefs bot",
			userAgent:  "Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)",
			deviceType: "desktop",
			browser:    "[Bot] Ahrefs",
			device:     "unknown",
			os:         "unknown",
			isBot:      true,
		},
		// Generic crawler
		{
			name:       "Generic crawler",
			userAgent:  "Mozilla/5.0 (compatible; MyCrawler/1.0)",
			deviceType: "desktop",
			browser:    "[Bot] Other",
			device:     "unknown",
			os:         "unknown",
			isBot:      true,
		},
		// Edge cases
		{
			name:       "Empty string",
			userAgent:  "",
			deviceType: "desktop",
			browser:    "unknown",
			device:     "unknown",
			os:         "unknown",
			isBot:      true,
		},
		{
			name:       "Random gibberish",
			userAgent:  "xyzzy plugh 12345",
			deviceType: "desktop",
			browser:    "unknown",
			device:     "unknown",
			os:         "unknown",
			isBot:      true,
		},
		{
			name:       "Very long string",
			userAgent:  strings.Repeat("A", 10000),
			deviceType: "desktop",
			browser:    "unknown",
			device:     "unknown",
			os:         "unknown",
			isBot:      true,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ua := Parse(tc.userAgent)
			if ua.DeviceType() != tc.deviceType {
				t.Errorf("expected device type %q, but got %q", tc.deviceType, ua.DeviceType())
			}

			if ua.Browser() != tc.browser {
				t.Errorf("expected browser %q, but got %q", tc.browser, ua.Browser())
			}

			if ua.Device() != tc.device {
				t.Errorf("expected device %q, but got %q", tc.device, ua.Device())
			}

			if ua.OperatingSystem() != tc.os {
				t.Errorf("expected OS %q, but got %q", tc.os, ua.OperatingSystem())
			}

			if ua.IsBot(true) != tc.isBot {
				t.Errorf("expected isBot %v, but got %v", tc.isBot, ua.IsBot(true))
			}
		})
	}
}

func TestHelperMethods(t *testing.T) {
	t.Parallel()

	t.Run("desktop Chrome on Windows", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		if !ua.IsDesktop() {
			t.Error("expected IsDesktop() to be true")
		}

		if ua.IsMobile() {
			t.Error("expected IsMobile() to be false")
		}

		if ua.IsTablet() {
			t.Error("expected IsTablet() to be false")
		}

		if !ua.IsWindows() {
			t.Error("expected IsWindows() to be true")
		}

		if ua.IsMacOS() {
			t.Error("expected IsMacOS() to be false")
		}

		if ua.IsLinux() {
			t.Error("expected IsLinux() to be false")
		}

		if ua.IsAndroid() {
			t.Error("expected IsAndroid() to be false")
		}

		if ua.IsIOS() {
			t.Error("expected IsIOS() to be false")
		}

		if !ua.IsValid() {
			t.Error("expected IsValid() to be true")
		}

		if ua.IsBot(true) {
			t.Error("expected IsBot(true) to be false")
		}

		if !ua.IsBrowserValid() {
			t.Error("expected IsBrowserValid() to be true")
		}

		if !ua.IsOperatingSystemValid() {
			t.Error("expected IsOperatingSystemValid() to be true")
		}

		if !ua.IsDeviceValid() {
			t.Error("expected IsDeviceValid() to be true")
		}
	})

	t.Run("mobile Safari on iPhone", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1")

		if ua.IsDesktop() {
			t.Error("expected IsDesktop() to be false")
		}

		if !ua.IsMobile() {
			t.Error("expected IsMobile() to be true")
		}

		if !ua.IsIOS() {
			t.Error("expected IsIOS() to be true")
		}

		if ua.IsAndroid() {
			t.Error("expected IsAndroid() to be false")
		}
	})

	t.Run("tablet iPad", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (iPad; CPU OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1")

		if !ua.IsTablet() {
			t.Error("expected IsTablet() to be true")
		}

		if ua.IsDesktop() {
			t.Error("expected IsDesktop() to be false")
		}

		if !ua.IsIOS() {
			t.Error("expected IsIOS() to be true")
		}
	})

	t.Run("Android phone", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.144 Mobile Safari/537.36")

		if !ua.IsMobile() {
			t.Error("expected IsMobile() to be true")
		}

		if !ua.IsAndroid() {
			t.Error("expected IsAndroid() to be true")
		}
	})

	t.Run("macOS desktop", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 14_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15")

		if !ua.IsMacOS() {
			t.Error("expected IsMacOS() to be true")
		}

		if !ua.IsDesktop() {
			t.Error("expected IsDesktop() to be true")
		}
	})

	t.Run("Linux desktop", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		if !ua.IsLinux() {
			t.Error("expected IsLinux() to be true")
		}
	})

	t.Run("bot detection", func(t *testing.T) {
		t.Parallel()

		ua := Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

		if ua.IsValid() {
			t.Error("expected IsValid() to be false for bot")
		}

		if !ua.IsBot(true) {
			t.Error("expected IsBot(true) to be true")
		}

		if !ua.IsBot(false) {
			t.Error("expected IsBot(false) to be true")
		}
	})

	t.Run("UserAgent returns original string", func(t *testing.T) {
		t.Parallel()

		input := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0"
		ua := Parse(input)

		if ua.UserAgent() != input {
			t.Errorf("expected UserAgent() to return %q, got %q", input, ua.UserAgent())
		}
	})
}

func TestWOW64BugFix(t *testing.T) {
	t.Parallel()

	// WOW64 should NOT cause false Windows 8 detection on Windows 10 64-bit
	ua := Parse("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	if ua.Device() != "Windows 10" {
		t.Errorf("expected device %q, but got %q (WOW64 false positive)", "Windows 10", ua.Device())
	}
}
