package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	arg := httprouter.Param{
		Key:   "arg_key",
		Value: "arg_value",
	}
	args := httprouter.Params{arg}

	v := url.Values{}
	v.Set("param_key", "param_value")
	req.PostForm = v

	// new context
	ctx := NewContext(rec, req, Args{args})
	assert.Equal(t, &Request{req}, ctx.Req())
	assert.Equal(t, NewResponse(), ctx.Res())
	assert.Equal(t, NewView(), ctx.View())
	assert.Equal(t, Args{args}, ctx.Args())
	assert.Equal(t, "arg_value", ctx.Arg("arg_key"))
	assert.Equal(t, Params{"param_key": []string{"param_value"}}, ctx.Params())
	assert.Equal(t, "param_value", ctx.ParamSingle("param_key"))
}
