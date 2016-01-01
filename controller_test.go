package ksatriya

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	// new controller
	c := NewController()
	assert.Equal(t, []*Handler{}, c.Routes())

	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	ctx := newTestContext(t)

	for i, method := range methods {
		path := fmt.Sprintf("/%s", strings.ToLower(method))

		called := false
		hf := func(ctx Ctx) {
			called = true
		}

		// add route
		switch method {
		case "GET":
			c.GET(path, hf)
		case "POST":
			c.POST(path, hf)
		case "PUT":
			c.PUT(path, hf)
		case "PATCH":
			c.PATCH(path, hf)
		case "DELETE":
			c.DELETE(path, hf)
		}
		assert.Len(t, c.Routes(), i+1)

		// test handler
		h := c.Routes()[i]
		h.HandlerFunc()(ctx)
		assert.Equal(t, method, h.Method())
		assert.Equal(t, path, h.Path())
		assert.True(t, called)
	}
}
