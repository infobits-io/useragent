# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

Un package Go léger pour analyser les chaînes d'agent utilisateur. Détecte les navigateurs, systèmes d'exploitation, appareils, types d'appareils et bots. Aucune dépendance.

## Installation

```go
go get github.com/infobits-io/useragent
```

## Utilisation

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

Analyse une chaîne d'agent utilisateur et renvoie un `*UserAgent` avec les informations détectées du navigateur, du système d'exploitation, de l'appareil et du bot.

### Méthodes de `UserAgent`

| Méthode | Type de retour | Description |
|---|---|---|
| `UserAgent()` | `string` | Chaîne d'agent utilisateur d'origine |
| `Browser()` | `string` | Nom du navigateur détecté |
| `OperatingSystem()` | `string` | Système d'exploitation détecté |
| `Device()` | `string` | Appareil détecté |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` ou `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Si l'agent utilisateur est un bot |
| `IsValid()` | `bool` | Si le navigateur, le SE et l'appareil sont tous reconnus |
| `IsBrowserValid()` | `bool` | Si le navigateur est reconnu |
| `IsOperatingSystemValid()` | `bool` | Si le système d'exploitation est reconnu |
| `IsDeviceValid()` | `bool` | Si l'appareil est reconnu |
| `IsMobile()` | `bool` | Si le type d'appareil est mobile |
| `IsTablet()` | `bool` | Si le type d'appareil est tablette |
| `IsDesktop()` | `bool` | Si le type d'appareil est bureau |
| `IsWindows()` | `bool` | Si le système d'exploitation est Windows |
| `IsLinux()` | `bool` | Si le système d'exploitation est Linux |
| `IsMacOS()` | `bool` | Si le système d'exploitation est macOS |
| `IsAndroid()` | `bool` | Si le système d'exploitation est Android |
| `IsIOS()` | `bool` | Si le système d'exploitation est iOS |

## Navigateurs pris en charge

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer et plus encore. Voir [`useragent.go`](useragent.go) pour la liste complète.

## Bots pris en charge

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush et plus encore. Voir [`useragent.go`](useragent.go) pour la liste complète.

## Licence

[BSD 3-Clause](LICENSE)
