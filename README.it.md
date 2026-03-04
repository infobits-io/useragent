# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

Un pacchetto Go leggero per analizzare le stringhe user agent. Rileva browser, sistemi operativi, dispositivi, tipi di dispositivi e bot. Nessuna dipendenza.

## Installazione

```go
go get github.com/infobits-io/useragent
```

## Utilizzo

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

Analizza una stringa user agent e restituisce un `*UserAgent` con le informazioni rilevate su browser, sistema operativo, dispositivo e bot.

### Metodi di `UserAgent`

| Metodo | Tipo di ritorno | Descrizione |
|---|---|---|
| `UserAgent()` | `string` | Stringa user agent originale |
| `Browser()` | `string` | Nome del browser rilevato |
| `OperatingSystem()` | `string` | Sistema operativo rilevato |
| `Device()` | `string` | Dispositivo rilevato |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` o `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Se lo user agent è un bot |
| `IsValid()` | `bool` | Se browser, SO e dispositivo sono tutti riconosciuti |
| `IsBrowserValid()` | `bool` | Se il browser è riconosciuto |
| `IsOperatingSystemValid()` | `bool` | Se il sistema operativo è riconosciuto |
| `IsDeviceValid()` | `bool` | Se il dispositivo è riconosciuto |
| `IsMobile()` | `bool` | Se il tipo di dispositivo è mobile |
| `IsTablet()` | `bool` | Se il tipo di dispositivo è tablet |
| `IsDesktop()` | `bool` | Se il tipo di dispositivo è desktop |
| `IsWindows()` | `bool` | Se il sistema operativo è Windows |
| `IsLinux()` | `bool` | Se il sistema operativo è Linux |
| `IsMacOS()` | `bool` | Se il sistema operativo è macOS |
| `IsAndroid()` | `bool` | Se il sistema operativo è Android |
| `IsIOS()` | `bool` | Se il sistema operativo è iOS |

## Browser supportati

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer e altri. Vedi [`useragent.go`](useragent.go) per l'elenco completo.

## Bot supportati

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush e altri. Vedi [`useragent.go`](useragent.go) per l'elenco completo.

## Licenza

[BSD 3-Clause](LICENSE)
