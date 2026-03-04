# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

En letvægts Go-pakke til at parse user agent-strenge. Detekterer browsere, operativsystemer, enheder, enhedstyper og bots. Ingen afhængigheder.

## Installation

```go
go get github.com/infobits-io/useragent
```

## Brug

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

Parser en user agent-streng og returnerer en `*UserAgent` med detekteret browser, OS, enhed og bot-information.

### `UserAgent`-metoder

| Metode | Returtype | Beskrivelse |
|---|---|---|
| `UserAgent()` | `string` | Oprindelig user agent-streng |
| `Browser()` | `string` | Detekteret browsernavn |
| `OperatingSystem()` | `string` | Detekteret operativsystem |
| `Device()` | `string` | Detekteret enhed |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` eller `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Om user agenten er en bot |
| `IsValid()` | `bool` | Om browser, OS og enhed alle er genkendt |
| `IsBrowserValid()` | `bool` | Om browseren er genkendt |
| `IsOperatingSystemValid()` | `bool` | Om operativsystemet er genkendt |
| `IsDeviceValid()` | `bool` | Om enheden er genkendt |
| `IsMobile()` | `bool` | Om enhedstypen er mobil |
| `IsTablet()` | `bool` | Om enhedstypen er tablet |
| `IsDesktop()` | `bool` | Om enhedstypen er desktop |
| `IsWindows()` | `bool` | Om operativsystemet er Windows |
| `IsLinux()` | `bool` | Om operativsystemet er Linux |
| `IsMacOS()` | `bool` | Om operativsystemet er macOS |
| `IsAndroid()` | `bool` | Om operativsystemet er Android |
| `IsIOS()` | `bool` | Om operativsystemet er iOS |

## Understottede browsere

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer og flere. Se [`useragent.go`](useragent.go) for den fulde liste.

## Understottede bots

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush og flere. Se [`useragent.go`](useragent.go) for den fulde liste.

## Licens

[BSD 3-Clause](LICENSE)
