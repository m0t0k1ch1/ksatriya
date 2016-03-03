package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func newTestContext() Ctx {
	rec := httptest.NewRecorder()
	req := &http.Request{}

	return NewContext(rec, req, Args{})
}

func TestContext(t *testing.T) {
	rec := httptest.NewRecorder()
	req := &http.Request{}

	arg := httprouter.Param{
		Key:   "arg_key",
		Value: "arg_value",
	}
	args := httprouter.Params{arg}

	v := url.Values{}
	v.Set("param_key", "param_value")
	req.PostForm = v

	ctx := NewContext(rec, req, Args{args})
	assert.Equal(t, &Request{req}, ctx.Req())
	assert.Equal(t, NewResponse(), ctx.Res())
	assert.Equal(t, NewView(), ctx.View())
	assert.Equal(t, Args{args}, ctx.Args())
	assert.Equal(t, "arg_value", ctx.Arg("arg_key"))
	assert.Equal(t, Params{"param_key": []string{"param_value"}}, ctx.Params())
	assert.Equal(t, "param_value", ctx.ParamSingle("param_key"))
}

func TestContext_RenderText(t *testing.T) {
	ctx := newTestContext()

	ctx.RenderText(http.StatusOK, "text")
	assert.Equal(t, http.StatusOK, ctx.Res().StatusCode())
	assert.Equal(t, "text/plain", ctx.Res().Header().Get("Content-Type"))
	assert.Equal(t, NewTextRenderer("text"), ctx.View().Renderer())

	rec := httptest.NewRecorder()

	ctx.WriteResponse(rec)
	assert.Equal(t, "text", rec.Body.String())
}

func TestContext_RenderJSON(t *testing.T) {
	ctx := newTestContext()

	j := map[string]string{"key": "value"}

	ctx.RenderJSON(http.StatusOK, j)
	assert.Equal(t, http.StatusOK, ctx.Res().StatusCode())
	assert.Equal(t, "application/json", ctx.Res().Header().Get("Content-Type"))
	assert.Equal(t, NewJSONRenderer(j), ctx.View().Renderer())

	rec := httptest.NewRecorder()

	ctx.WriteResponse(rec)
	assert.Equal(t, "{\"key\":\"value\"}", rec.Body.String())
}

func TestContext_Redirect(t *testing.T) {
	ctx := newTestContext()

	ctx.Redirect("/redirect")
	assert.Equal(t, http.StatusFound, ctx.Res().StatusCode())
	assert.Equal(t, "/redirect", ctx.Res().Header().Get("Location"))
	assert.Equal(t, NewTextRenderer(""), ctx.View().Renderer())

	rec := httptest.NewRecorder()

	ctx.WriteResponse(rec)
	assert.Empty(t, rec.Body.String())
}
