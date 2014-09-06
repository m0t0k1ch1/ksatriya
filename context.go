package ksatriya

import "net/http"

type Context interface {
	Request() *http.Request
	Response() *Response
	Params() Params
	Param(name string) string
	Renderer() *Renderer
	RenderArgs() RenderArgs
	SetRenderArg(key string, value interface{})
	SetStatusCode(statusCode int)
	SetTmplDir(tmplDir string)
	SetBaseTmplPath(baseTmplPath string)
	RenderText(statusCode int, text string)
	RenderJSON(statusCode int, data interface{})
	RenderHTML(statusCode int, tmplPath string, renderArgs RenderArgs)
	Write(w http.ResponseWriter)
}

type Ctx struct {
	request    *http.Request
	response   *Response
	params     Params
	renderer   *Renderer
	renderArgs RenderArgs
}

func NewContext(req *http.Request, params Params, r *Renderer) Context {
	req.ParseForm()
	return &Ctx{
		request: req,
		response: &Response{
			StatusCode: 200,
			Header:     http.Header{},
		},
		params:     params,
		renderer:   r,
		renderArgs: RenderArgs{},
	}
}

func (ctx *Ctx) Request() *http.Request {
	return ctx.request
}

func (ctx *Ctx) Response() *Response {
	return ctx.response
}

func (ctx *Ctx) Params() Params {
	return ctx.params
}

func (ctx *Ctx) Param(name string) string {
	return ctx.params.ByName(name)
}

func (ctx *Ctx) Renderer() *Renderer {
	return ctx.renderer
}

func (ctx *Ctx) RenderArgs() RenderArgs {
	return ctx.renderArgs
}

func (ctx *Ctx) SetRenderArg(key string, value interface{}) {
	ctx.renderArgs[key] = value
}

func (ctx *Ctx) SetStatusCode(statusCode int) {
	ctx.response.StatusCode = statusCode
}

func (ctx *Ctx) SetTmplDir(tmplDir string) {
	ctx.renderer.TmplDir = tmplDir
}

func (ctx *Ctx) SetBaseTmplPath(baseTmplPath string) {
	ctx.renderer.BaseTmplPath = baseTmplPath
}

func (ctx *Ctx) RenderText(statusCode int, text string) {
	ctx.SetStatusCode(statusCode)
	ctx.response.Result = ctx.renderer.RenderText(text)
}

func (ctx *Ctx) RenderJSON(statusCode int, data interface{}) {
	ctx.SetStatusCode(statusCode)
	ctx.response.Result = ctx.renderer.RenderJSON(data)
}

func (ctx *Ctx) RenderHTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	ctx.SetStatusCode(statusCode)
	for k, v := range renderArgs {
		ctx.renderArgs[k] = v
	}
	ctx.response.Result = ctx.renderer.RenderHTML(tmplPath)
}

func (ctx *Ctx) Write(w http.ResponseWriter) {
	ctx.response.Write(ctx, w)
}
