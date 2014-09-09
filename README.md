ksatriya
========

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/ksatriya?status.svg)](https://godoc.org/github.com/m0t0k1ch1/ksatriya)

a tiny web application framework for golang

**NOTICE: ksatriya is still under development, so API might be changed in future.**

``` sh
$ go get github.com/m0t0k1ch1/ksatriya
```

## Examples

ref. https://github.com/m0t0k1ch1/ksatriya-sample

### Basic

``` go
package main

import (
    "fmt"
    "net/http"

    "github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
    *ksatriya.Controller
}

func NewController() *Controller {
    c := &Controller{ksatriya.NewController()}
    c.GET("/", c.Index)
    c.GET("/user/:name", c.User)
    return c
}

func (c *Controller) Index(ctx *ksatriya.Context) {
    ctx.Text(http.StatusOK, "index")
}

func (c *Controller) User(ctx *ksatriya.Context) {
    name := ctx.Param("name")
    ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.RegisterController(NewController())
    k.Run(":8080")
}
```

### With [negroni](https://github.com/codegangsta/negroni) and [go-server-starter-listner](https://github.com/lestrrat/go-server-starter-listener)

``` go
package main

import (
    "fmt"
    "net"
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/lestrrat/go-server-starter-listener"
    "github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
    *ksatriya.Controller
}

func NewController() *Controller {
    c := &Controller{ksatriya.NewController()}
    c.GET("/", c.Index)
    c.GET("/user/:name", c.User)
    return c
}

func (c *Controller) Index(ctx *ksatriya.Context) {
    ctx.Text(http.StatusOK, "index")
}

func (c *Controller) User(ctx *ksatriya.Context) {
    name := ctx.Param("name")
    ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.RegisterController(NewController())

    n := negroni.Classic()
    n.UseHandler(k)

    listener, _ := ss.NewListener()
    if listener == nil {
        var err error
        listener, err = net.Listen("tcp", ":8080")
        if err != nil {
            panic(err)
        }
    }

    server := &http.Server{Handler: n}
    server.Serve(listener)
}
```

## Dependencies

* routing - https://github.com/julienschmidt/httprouter
