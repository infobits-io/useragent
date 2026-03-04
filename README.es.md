# useragent

[![CI](https://github.com/infobits-io/useragent/actions/workflows/ci.yml/badge.svg)](https://github.com/infobits-io/useragent/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/infobits-io/useragent.svg)](https://pkg.go.dev/github.com/infobits-io/useragent)

Un paquete ligero de Go para analizar cadenas de agente de usuario. Detecta navegadores, sistemas operativos, dispositivos, tipos de dispositivos y bots. Sin dependencias.

## Instalación

```go
go get github.com/infobits-io/useragent
```

## Uso

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

Analiza una cadena de agente de usuario y devuelve un `*UserAgent` con información detectada del navegador, sistema operativo, dispositivo y bot.

### Métodos de `UserAgent`

| Método | Tipo de retorno | Descripción |
|---|---|---|
| `UserAgent()` | `string` | Cadena de agente de usuario original |
| `Browser()` | `string` | Nombre del navegador detectado |
| `OperatingSystem()` | `string` | Sistema operativo detectado |
| `Device()` | `string` | Dispositivo detectado |
| `DeviceType()` | `string` | `"desktop"`, `"mobile"` o `"tablet"` |
| `IsBot(includeBrowser bool)` | `bool` | Si el agente de usuario es un bot |
| `IsValid()` | `bool` | Si el navegador, SO y dispositivo son todos reconocidos |
| `IsBrowserValid()` | `bool` | Si el navegador es reconocido |
| `IsOperatingSystemValid()` | `bool` | Si el sistema operativo es reconocido |
| `IsDeviceValid()` | `bool` | Si el dispositivo es reconocido |
| `IsMobile()` | `bool` | Si el tipo de dispositivo es móvil |
| `IsTablet()` | `bool` | Si el tipo de dispositivo es tablet |
| `IsDesktop()` | `bool` | Si el tipo de dispositivo es escritorio |
| `IsWindows()` | `bool` | Si el sistema operativo es Windows |
| `IsLinux()` | `bool` | Si el sistema operativo es Linux |
| `IsMacOS()` | `bool` | Si el sistema operativo es macOS |
| `IsAndroid()` | `bool` | Si el sistema operativo es Android |
| `IsIOS()` | `bool` | Si el sistema operativo es iOS |

## Navegadores compatibles

Chrome, Safari, Firefox, Edge, Opera, Brave, DuckDuckGo, Samsung Internet, UC Browser, Vivaldi, Tor Browser, Internet Explorer y más. Consulta [`useragent.go`](useragent.go) para la lista completa.

## Bots compatibles

Googlebot, Bingbot, Baidu, Yandex, DuckDuckBot, Facebook, Twitter, LinkedIn, Instagram, ChatGPT, OpenAI, Ahrefs, SEMRush y más. Consulta [`useragent.go`](useragent.go) para la lista completa.

## Licencia

[BSD 3-Clause](LICENSE)
