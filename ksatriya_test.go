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

	h := k.ServeHTTP

	assert.HTTPSuccess(t, h, "GET", "/user/m0t0k1ch1", nil)
	assert.HTTPBodyContains(t, h, "GET", "/user/m0t0k1ch1", nil, "your name is m0t0k1ch1")
}

func TestKsatriya_ServeHTTP_redirection(t *testing.T) {
	k := New()
	k.GET("/", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "")
	})
	k.GET("/redirect", func(ctx Ctx) {
		ctx.Redirect("/")
	})

	h := k.ServeHTTP

	assert.HTTPRedirect(t, h, "GET", "/redirect", nil)
}

func TestKsatriya_ServeHTTP_withController(t *testing.T) {
	c := NewController()
	c.GET("/ping", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "pong")
	})

	k := New()
	k.RegisterController(c)

	h := k.ServeHTTP

	assert.HTTPSuccess(t, h, "GET", "/ping", nil)
	assert.HTTPBodyContains(t, h, "GET", "/ping", nil, "pong")
}
