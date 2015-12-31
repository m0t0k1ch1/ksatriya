package ksatriya

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKsatriya_ServeHTTP(t *testing.T) {
	k := New()
	k.GET("/user/:name", func(ctx Ctx) {
		name := ctx.Arg("name")
		res := fmt.Sprintf("your name is %s", name)
		ctx.RenderText(http.StatusOK, res)
	})
	handler := k.ServeHTTP

	assert.HTTPSuccess(t, handler, "GET", "/user/m0t0k1ch1", nil)
	assert.HTTPBodyContains(t, handler, "GET", "/user/m0t0k1ch1", nil, "your name is m0t0k1ch1")
}

func TestKsatriya_ServeHTTP_redirection(t *testing.T) {
	k := New()
	k.GET("/", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "")
	})
	k.GET("/redirect", func(ctx Ctx) {
		ctx.Redirect("/")
	})
	handler := k.ServeHTTP

	assert.HTTPRedirect(t, handler, "GET", "/redirect", nil)
}
