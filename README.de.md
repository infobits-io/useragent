# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

Ein leichtgewichtiges Go-Paket zum Parsen von User-Agent-Strings. Erkennt Browser, Betriebssysteme, Geräte, Gerätetypen und Bots. Keine Abhängigkeiten.

## Installation

```go
go get github.com/infobits-io/useragent
```

## Verwendung

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

Parst einen User-Agent-String und gibt ein `*UserAgent` mit erkanntem Browser, Betriebssystem, Gerät und Bot-Informationen zurück.

### `UserAgent`-Methoden

| Methode | Rückgabetyp | Beschreibung |
|---|---|---|
| `UserAgent()` | `string` | Ursprünglicher User-Agent-String |
| `Browser()` | `string` | Erkannter Browsername |
| `OperatingSystem()` | `string` | Erkanntes Betriebssystem |
| `Device()` | `string` | Erkanntes Gerät |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` oder `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Ob der User-Agent ein Bot ist |
| `IsValid()` | `bool` | Ob Browser, Betriebssystem und Gerät alle erkannt wurden |
| `IsBrowserValid()` | `bool` | Ob der Browser erkannt wurde |
| `IsOperatingSystemValid()` | `bool` | Ob das Betriebssystem erkannt wurde |
| `IsDeviceValid()` | `bool` | Ob das Gerät erkannt wurde |
| `IsMobile()` | `bool` | Ob der Gerätetyp mobil ist |
| `IsTablet()` | `bool` | Ob der Gerätetyp Tablet ist |
| `IsDesktop()` | `bool` | Ob der Gerätetyp Desktop ist |
| `IsWindows()` | `bool` | Ob das Betriebssystem Windows ist |
| `IsLinux()` | `bool` | Ob das Betriebssystem Linux ist |
| `IsMacOS()` | `bool` | Ob das Betriebssystem macOS ist |
| `IsAndroid()` | `bool` | Ob das Betriebssystem Android ist |
| `IsIOS()` | `bool` | Ob das Betriebssystem iOS ist |

## Unterstützte Browser

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer und weitere. Siehe [`useragent.go`](useragent.go) für die vollständige Liste.

## Unterstützte Bots

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush und weitere. Siehe [`useragent.go`](useragent.go) für die vollständige Liste.

## Lizenz

[BSD 3-Clause](LICENSE)
