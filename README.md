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

func Index(ctx *ksatriya.Context) {
	ctx.Text(http.StatusOK, "index")
}

func User(ctx *ksatriya.Context) {
	name := ctx.Arg("name")
	ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
	k := ksatriya.New()
	k.GET("/", Index)
	k.GET("/user/:name", User)
	k.Run(":8080")
}
```

### Use controller

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
    name := ctx.Arg("name")
    ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.RegisterController(NewController())
    k.Run(":8080")
}
```

### With [negroni](https://github.com/codegangsta/negroni) and [go-server-starter](https://github.com/lestrrat/go-server-starter)

``` go
package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/lestrrat/go-server-starter/listener"
	"github.com/m0t0k1ch1/ksatriya"
)

func Index(ctx *ksatriya.Context) {
	ctx.Text(http.StatusOK, "index")
}

func User(ctx *ksatriya.Context) {
	name := ctx.Arg("name")
	ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
	k := ksatriya.New()
	k.GET("/", Index)
	k.GET("/user/:name", User)

	n := negroni.Classic()
	n.UseHandler(k)

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}
	var l net.Listener
	if len(listeners) == 0 {
		l, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}
	} else {
		l = listeners[0]
	}

	server := &http.Server{Handler: n}
	server.Serve(l)
}
```

## Dependencies

* routing - https://github.com/julienschmidt/httprouter
