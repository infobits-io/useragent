# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

A lightweight Go package for parsing user agent strings. Detects browsers, operating systems, devices, device types, and bots. Zero dependencies.

## Installation

```go
go get github.com/infobits-io/useragent
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/infobits-io/useragent"
)

func main() {
    ua := useragent.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

    fmt.Println(ua.Browser())         // Chrome
    fmt.Println(ua.OperatingSystem()) // windows
    fmt.Println(ua.Device())          // Windows 10
    fmt.Println(ua.DeviceType())      // desktop
    fmt.Println(ua.IsBot(true))       // false
}
```

## API

### `Parse(userAgent string) *UserAgent`

Parses a user agent string and returns a `*UserAgent` with detected browser, OS, device, and bot information.

### `UserAgent` methods

| Method | Return type | Description |
|---|---|---|
| `UserAgent()` | `string` | Original user agent string |
| `Browser()` | `string` | Detected browser name |
| `OperatingSystem()` | `string` | Detected operating system |
| `Device()` | `string` | Detected device |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"`, or `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Whether the user agent is a bot |
| `IsValid()` | `bool` | Whether browser, OS, and device are all recognized |
| `IsBrowserValid()` | `bool` | Whether the browser is recognized |
| `IsOperatingSystemValid()` | `bool` | Whether the OS is recognized |
| `IsDeviceValid()` | `bool` | Whether the device is recognized |
| `IsMobile()` | `bool` | Whether the device type is mobile |
| `IsTablet()` | `bool` | Whether the device type is tablet |
| `IsDesktop()` | `bool` | Whether the device type is desktop |
| `IsWindows()` | `bool` | Whether the OS is Windows |
| `IsLinux()` | `bool` | Whether the OS is Linux |
| `IsMacOS()` | `bool` | Whether the OS is macOS |
| `IsAndroid()` | `bool` | Whether the OS is Android |
| `IsIOS()` | `bool` | Whether the OS is iOS |

## Supported Browsers

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer, and more. See [`useragent.go`](useragent.go) for the full list.

## Supported Bots

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush, and more. See [`useragent.go`](useragent.go) for the full list.

## License

[BSD 3-Clause](LICENSE)
