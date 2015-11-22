ksatriya
========

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/ksatriya?status.svg)](https://godoc.org/github.com/m0t0k1ch1/ksatriya)

a tiny web application framework for golang

* **ksatriya** do not support HTML rendering
* **ksatriya** is still under development, so API might be changed in future

``` sh
$ go get github.com/m0t0k1ch1/ksatriya
```

## Examples

ref. https://github.com/m0t0k1ch1/ksatriya-sample

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

func Ping (ctx ksatriya.Ctx) {
	ctx.Text(http.StatusOK, "pong")
}

func main() {
	k := ksatriya.New()
	k.GET("/ping", Ping)
	k.Run(":8080")
}
```

## Dependencies

* routing - https://github.com/julienschmidt/httprouter
