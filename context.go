package ksatriya

import "net/http"

type Context struct {
	Request  *http.Request
	Writer   http.ResponseWriter
	Params   Params
	Renderer *Renderer
}

func NewContext(w http.ResponseWriter, req *http.Request, params Params, r *Renderer) *Context {
	req.ParseForm()
	return &Context{
		Request:  req,
		Writer:   w,
		Params:   params,
		Renderer: r,
	}
}

func (c *Context) Param(name string) string {
	return c.Params.ByName(name)
}

func (c *Context) RenderHTML(status int, name string, data RenderData, layout ...string) {
	c.Renderer.RenderHTML(c.Writer, status, name, data, layout)
}

func (c *Context) RenderJSON(status int, v interface{}) {
	c.Renderer.RenderJSON(c.Writer, status, v)
}
