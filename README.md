ksatriya
========

a tiny web application framework for golang

**NOTICE: ksatriya is still under development, so API might be changed in future.**

## Getting Started

### Get ksatriya

``` sh
$ go get github.com/m0t0k1ch1/ksatriya
```

### Example

#### Basic

``` go
package main

import (
    "fmt"
    "net/http"

    "github.com/m0t0k1ch1/ksatriya"
)

func Index(c *ksatriya.Context) {
    c.RenderText(http.StatusOK, "index")
}

func User(c *ksatriya.Context) {
    name := c.Param("name")
    c.RenderText(http.StatusOK, fmt.Sprintf("hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.GET("/", Index)
    k.GET("/user/:name", User)
    k.Run(":8080")
}
```

#### With [negroni](https://github.com/codegangsta/negroni) and [go-server-starter-listner](https://github.com/lestrrat/go-server-starter-listener)

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

func Index(c *ksatriya.Context) {
    c.RenderText(http.StatusOK, "index")
}

func User(c *ksatriya.Context) {
    name := c.Param("name")
    c.RenderText(http.StatusOK, fmt.Sprintf("hello %s!", name))
}

func main() {
    k := ksatriya.New()
    k.GET("/", Index)
    k.GET("/user/:name", User)

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
* rendering - https://github.com/unrolled/render
