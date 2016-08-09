ksatriya
========

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/ksatriya?status.svg)](https://godoc.org/github.com/m0t0k1ch1/ksatriya) [![wercker status](https://app.wercker.com/status/31a22e629614ce105b38f9c4cb3326f5/s/master "wercker status")](https://app.wercker.com/project/bykey/31a22e629614ce105b38f9c4cb3326f5)

a tiny web application framework for golang

* **ksatriya** does not support HTML rendering
* **ksatriya** is still under development, so API might be changed in future

``` sh
$ go get github.com/m0t0k1ch1/ksatriya
```

## Example

ref. https://github.com/m0t0k1ch1/ksatriya-sample

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/m0t0k1ch1/ksatriya"
)

func Ping (ctx ksatriya.Ctx) {
	ctx.RenderText(http.StatusOK, "pong")
}

func main() {
	k := ksatriya.New()
	k.GET("/ping", Ping)
	k.Run(":8080")
}
```

## Dependencies

* routing - https://github.com/julienschmidt/httprouter
* test - https://github.com/stretchr/testify
