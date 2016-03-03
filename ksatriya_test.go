package ksatriya

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyTestContext struct {
	Ctx
	testField string
}

func TestKsatriya(t *testing.T) {
	rec := httptest.NewRecorder()
	req := &http.Request{}

	// new ksatriya
	k := New()
	assert.Equal(t, NewRouter(), k.Router())
	assert.Equal(t, NewRoot(), k.Root())
	assert.Equal(t, NewContext(rec, req, Args{}), k.ctxBuilder(rec, req, Args{}))

	// add routes
	k.GET("/get", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "get")
	})
	k.POST("/post", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "post")
	})
	k.PUT("/put", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "put")
	})
	k.PATCH("/patch", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "patch")
	})
	k.DELETE("/delete", func(ctx Ctx) {
		ctx.RenderText(http.StatusOK, "delete")
	})

	h := k.ServeHTTP

	// GET /get
	assert.HTTPSuccess(t, h, "GET", "/get", nil)
	assert.HTTPBodyContains(t, h, "GET", "/get", nil, "get")

	// POST /post
	assert.HTTPSuccess(t, h, "POST", "/post", nil)
	assert.HTTPBodyContains(t, h, "POST", "/post", nil, "post")

	// PUT /put
	assert.HTTPSuccess(t, h, "PUT", "/put", nil)
	assert.HTTPBodyContains(t, h, "PUT", "/put", nil, "put")

	// PATCH /patch
	assert.HTTPSuccess(t, h, "PATCH", "/patch", nil)
	assert.HTTPBodyContains(t, h, "PATCH", "/patch", nil, "patch")

	// DELETE /delete
	assert.HTTPSuccess(t, h, "DELETE", "/delete", nil)
	assert.HTTPBodyContains(t, h, "DELETE", "/delete", nil, "delete")
}

func TestKsatriya_redirection(t *testing.T) {
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

func TestKsatriya_errorHandling(t *testing.T) {
	k := New()
	k.GET("/", func(ctx Ctx) {
		ctx.RenderJSON(http.StatusOK, func() {
			return
		})
	})

	h := k.ServeHTTP

	assert.HTTPError(t, h, "GET", "/", nil)
}

func TestKsatriya_withArgs(t *testing.T) {
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

func TestKsatriya_withController(t *testing.T) {
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

func TestKsatriya_withCtxBuilder(t *testing.T) {
	k := New()
	k.SetCtxBuilder(func(w http.ResponseWriter, req *http.Request, args Args) Ctx {
		return &MyTestContext{
			Ctx:       NewContext(w, req, args),
			testField: "test_value",
		}
	})
	k.GET("/", func(ctx Ctx) {
		myctx := ctx.(*MyTestContext)
		ctx.RenderText(http.StatusOK, myctx.testField)
	})

	h := k.ServeHTTP

	assert.HTTPSuccess(t, h, "GET", "/", nil)
	assert.HTTPBodyContains(t, h, "GET", "/", nil, "test_value")
}
