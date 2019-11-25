# [gioui.org - GoDoc](https://godoc.org/gioui.org)

# Links

* [Gio - immediate mode GUI in Go](https://gioui.org/)

# Packages

* [app - GoDoc](https://godoc.org/gioui.org/app)

  Package app provides a platform-independent interface to operating system
  functionality for running graphical user interfaces. 

* [event - GoDoc](https://godoc.org/gioui.org/io/event)

  Package event contains the types for event handling. 

* [unit - GoDoc](https://godoc.org/gioui.org/unit)

  Package unit implements device independent units and values. 

* [layout - GoDoc](https://godoc.org/gioui.org/layout)

  Package layout implements layouts common to GUI programs. 

* [gofont - GoDoc](https://godoc.org/gioui.org/font/gofont)

  Package gofont registers the Go fonts in the font registry. 

* [widget - GoDoc](https://godoc.org/gioui.org/widget)

  Package widget implements state tracking and event handling of common user
  interface controls. To draw widgets, use a theme packages such as package
  gioui.org/widget/material.

* [material - GoDoc](https://godoc.org/gioui.org/widget/material)

  Package material implements the Material design.

# Overview

# Webassembly

```sh
go run gioui.org/cmd/gogio -target js gioui.org/example/gophers
go get github.com/shurcooL/goexec
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("gophers")))'
```
`
