# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

Een lichtgewicht Go-pakket voor het parsen van user-agentstrings. Detecteert browsers, besturingssystemen, apparaten, apparaattypen en bots. Geen afhankelijkheden.

## Installatie

```go
go get github.com/infobits-io/useragent
```

## Gebruik

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

Parst een user-agentstring en geeft een `*UserAgent` terug met gedetecteerde browser-, besturingssysteem-, apparaat- en botinformatie.

### `UserAgent`-methoden

| Methode | Retourtype | Beschrijving |
|---|---|---|
| `UserAgent()` | `string` | Oorspronkelijke user-agentstring |
| `Browser()` | `string` | Gedetecteerde browsernaam |
| `OperatingSystem()` | `string` | Gedetecteerd besturingssysteem |
| `Device()` | `string` | Gedetecteerd apparaat |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` of `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Of de user-agent een bot is |
| `IsValid()` | `bool` | Of browser, besturingssysteem en apparaat allemaal herkend zijn |
| `IsBrowserValid()` | `bool` | Of de browser herkend is |
| `IsOperatingSystemValid()` | `bool` | Of het besturingssysteem herkend is |
| `IsDeviceValid()` | `bool` | Of het apparaat herkend is |
| `IsMobile()` | `bool` | Of het apparaattype mobiel is |
| `IsTablet()` | `bool` | Of het apparaattype tablet is |
| `IsDesktop()` | `bool` | Of het apparaattype desktop is |
| `IsWindows()` | `bool` | Of het besturingssysteem Windows is |
| `IsLinux()` | `bool` | Of het besturingssysteem Linux is |
| `IsMacOS()` | `bool` | Of het besturingssysteem macOS is |
| `IsAndroid()` | `bool` | Of het besturingssysteem Android is |
| `IsIOS()` | `bool` | Of het besturingssysteem iOS is |

## Ondersteunde browsers

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer en meer. Zie [`useragent.go`](useragent.go) voor de volledige lijst.

## Ondersteunde bots

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush en meer. Zie [`useragent.go`](useragent.go) voor de volledige lijst.

## Licentie

[BSD 3-Clause](LICENSE)
