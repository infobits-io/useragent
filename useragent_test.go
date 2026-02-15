package useragent

import (
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
		{
			name:       "Chrome on Windows",
			userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
			deviceType: "desktop",
			browser:    "Chrome",
			device:     "Windows 10",
			os:         "windows",
			isBot:      false,
		},
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
			name:       "iPad with Safari",
			userAgent:  "Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
			deviceType: "tablet",
			browser:    "Safari",
			device:     "iPad",
			os:         "ios",
			isBot:      false,
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
