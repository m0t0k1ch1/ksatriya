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

func (ctx *Context) Param(name string) string {
	return ctx.Params.ByName(name)
}

func (ctx *Context) SetStatusCode(statusCode int) {
	ctx.Response.StatusCode = statusCode
}

func (ctx *Context) SetTmplDir(tmplDir string) {
	ctx.Renderer.TmplDir = tmplDir
}

func (ctx *Context) RenderText(text string) {
	ctx.Response.Result = ctx.Renderer.RenderText(text)
}

func (ctx *Context) RenderJSON(data interface{}) {
	ctx.Response.Result = ctx.Renderer.RenderJSON(data)
}

func (ctx *Context) RenderHTML(tmplPath string, renderData *RenderData) {
	ctx.Response.Result = ctx.Renderer.RenderHTML(tmplPath, renderData)
}
