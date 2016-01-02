package ksatriya

import (
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

	ctx := newTestContext()

	// GET /get
	hGET := c.Routes()[0]
	hGET.HandlerFunc()(ctx)
	assert.Equal(t, "GET", hGET.Method())
	assert.Equal(t, "/get", hGET.Path())
	assert.Equal(t, []string{"GET"}, calledMethods)

	// POST /post
	hPOST := c.Routes()[1]
	hPOST.HandlerFunc()(ctx)
	assert.Equal(t, "POST", hPOST.Method())
	assert.Equal(t, "/post", hPOST.Path())
	assert.Equal(t, []string{"GET", "POST"}, calledMethods)

	// PUT /put
	hPUT := c.Routes()[2]
	hPUT.HandlerFunc()(ctx)
	assert.Equal(t, "PUT", hPUT.Method())
	assert.Equal(t, "/put", hPUT.Path())
	assert.Equal(t, []string{"GET", "POST", "PUT"}, calledMethods)

	// PATCH /patch
	hPATCH := c.Routes()[3]
	hPATCH.HandlerFunc()(ctx)
	assert.Equal(t, "PATCH", hPATCH.Method())
	assert.Equal(t, "/patch", hPATCH.Path())
	assert.Equal(t, []string{"GET", "POST", "PUT", "PATCH"}, calledMethods)

	// DELETE /delete
	hDELETE := c.Routes()[4]
	hDELETE.HandlerFunc()(ctx)
	assert.Equal(t, "DELETE", hDELETE.Method())
	assert.Equal(t, "/delete", hDELETE.Path())
	assert.Equal(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, calledMethods)
}
