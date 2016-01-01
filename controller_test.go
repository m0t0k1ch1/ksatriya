package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	c := NewController()

	assert.Equal(t, []*Handler{}, c.Routes())

	called := false
	hf := func(ctx Ctx) {
		called = true
	}
	c.GET("/", hf)

	assert.Len(t, c.Routes(), 1)

	h := c.Routes()[0]

	assert.Equal(t, "GET", h.Method())
	assert.Equal(t, "/", h.Path())

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	ctx := NewContext(rec, req, Args{})
	h.HandlerFunc()(ctx)

	assert.True(t, called)
}
