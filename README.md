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

func Index(ctx ksatriya.Ctx) {
	ctx.Text(http.StatusOK, "index")
}

func User(ctx ksatriya.Ctx) {
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

func (c *Controller) Index(ctx ksatriya.Ctx) {
    ctx.Text(http.StatusOK, "index")
}

func (c *Controller) User(ctx ksatriya.Ctx) {
    name := ctx.Arg("name")
    ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.RegisterController(NewController())
    k.Run(":8080")
}
```

### Use original context

``` go
package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

type Context struct {
	ksatriya.Ctx
	client *http.Client
}

func NewContext(w http.ResponseWriter, req *http.Request, args ksatriya.Args) ksatriya.Ctx {
	return &Context{
		Ctx:    ksatriya.NewContext(w, req, args),
		client: &http.Client{},
	}
}

func convertContext(kctx ksatriya.Ctx) *Context {
	ctx, _ := kctx.(*Context)
	return ctx
}

func PingHandler(kctx ksatriya.Ctx) {
	Ping(convertContext(kctx))
}
func Ping(ctx *Context) {
	res, err := ctx.client.Get("http://127.0.0.1:8080/pong")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Text(http.StatusOK, string(body))
}

func PongHandler(kctx ksatriya.Ctx) {
	Pong(convertContext(kctx))
}
func Pong(ctx *Context) {
	ctx.Text(http.StatusOK, "pong")
}

func main() {
	k := ksatriya.New()
	k.SetCtxBuilder(NewContext)
	k.GET("/ping", PingHandler)
	k.GET("/pong", PongHandler)
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

func Index(ctx ksatriya.Ctx) {
	ctx.Text(http.StatusOK, "index")
}

func User(ctx ksatriya.Ctx) {
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
