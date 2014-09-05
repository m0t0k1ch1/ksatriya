package ksatriya

import "net/http"

type Context struct {
	Request  *http.Request
	Response *Response
	Params   Params
	Renderer *Renderer
	Stash    map[string]interface{}
}

func NewContext(req *http.Request, params Params, r *Renderer) *Context {
	req.ParseForm()
	return &Context{
		Request: req,
		Response: &Response{
			StatusCode: 200,
			Header:     http.Header{},
		},
		Params:   params,
		Renderer: r,
		Stash:    map[string]interface{}{},
	}
}

func (c *Context) Param(name string) string {
	return c.Params.ByName(name)
}

func (c *Context) SetStatusCode(statusCode int) {
	c.Response.StatusCode = statusCode
}

func (c *Context) RenderText(text string) Result {
	return c.Renderer.RenderText(text)
}

func (c *Context) RenderJSON(data interface{}) Result {
	return c.Renderer.RenderJSON(data)
}

func (c *Context) RenderHTML(tmplPath string, renderData *RenderData) Result {
	return c.Renderer.RenderHTML(tmplPath, renderData)
}
