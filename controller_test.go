package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	// new controller
	c := NewController()
	assert.Equal(t, []*Handler{}, c.Routes())

	calledMethods := []string{}

	// add routes
	c.GET("/get", func(ctx Ctx) {
		calledMethods = append(calledMethods, "GET")
	})
	c.POST("/post", func(ctx Ctx) {
		calledMethods = append(calledMethods, "POST")
	})
	c.PUT("/put", func(ctx Ctx) {
		calledMethods = append(calledMethods, "PUT")
	})
	c.PATCH("/patch", func(ctx Ctx) {
		calledMethods = append(calledMethods, "PATCH")
	})
	c.DELETE("/delete", func(ctx Ctx) {
		calledMethods = append(calledMethods, "DELETE")
	})
	assert.Len(t, c.Routes(), 5)

	// GET /get
	hGET := c.Routes()[0]
	assert.Equal(t, "GET", hGET.Method())
	assert.Equal(t, "/get", hGET.Path())

	// POST /post
	hPOST := c.Routes()[1]
	assert.Equal(t, "POST", hPOST.Method())
	assert.Equal(t, "/post", hPOST.Path())

	// PUT /put
	hPUT := c.Routes()[2]
	assert.Equal(t, "PUT", hPUT.Method())
	assert.Equal(t, "/put", hPUT.Path())

	// PATCH /patch
	hPATCH := c.Routes()[3]
	assert.Equal(t, "PATCH", hPATCH.Method())
	assert.Equal(t, "/patch", hPATCH.Path())

	// DELETE /delete
	hDELETE := c.Routes()[4]
	assert.Equal(t, "DELETE", hDELETE.Method())
	assert.Equal(t, "/delete", hDELETE.Path())

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	ctx := NewContext(rec, req, Args{})

	// call handler func
	hGET.HandlerFunc()(ctx)
	hPOST.HandlerFunc()(ctx)
	hPUT.HandlerFunc()(ctx)
	hPATCH.HandlerFunc()(ctx)
	hDELETE.HandlerFunc()(ctx)
	assert.Equal(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, calledMethods)
}
