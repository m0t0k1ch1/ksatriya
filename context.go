package ksatriya

import "net/http"

type Ctx interface {
	Req() *http.Request
	Res() Res
	Params() Params
	Param(name string) string
	View() ResultBuilder
	RenderArgs() RenderArgs

	SetRenderArg(key string, value interface{})

	Text(statusCode int, text string)
	JSON(statusCode int, data interface{})
	HTML(statusCode int, tmplPath string, renderArgs RenderArgs)
	Redirect(uri string)
}

type Context struct {
	request    *http.Request
	response   Res
	params     Params
	view       ResultBuilder
	renderArgs RenderArgs
}

func NewContext(w http.ResponseWriter, req *http.Request, params Params, v ResultBuilder) *Context {
	req.ParseForm()
	return &Context{
		request:    req,
		response:   NewResponse(w),
		params:     params,
		view:       v,
		renderArgs: RenderArgs{},
	}
}

func (ctx *Context) Req() *http.Request {
	return ctx.request
}

func (ctx *Context) Res() Res {
	return ctx.response
}

func (ctx *Context) Params() Params {
	return ctx.params
}

func (ctx *Context) Param(name string) string {
	return ctx.Params().ByName(name)
}

func (ctx *Context) View() ResultBuilder {
	return ctx.view
}

func (ctx *Context) RenderArgs() RenderArgs {
	return ctx.renderArgs
}

func (ctx *Context) SetRenderArg(key string, value interface{}) {
	ctx.renderArgs[key] = value
}

func (ctx *Context) Text(statusCode int, text string) {
	res := ctx.Res()
	res.SetStatusCode(statusCode)
	res.SetContentType("text/plain")
	res.SetResult(ctx.View().Text(text))
}

func (ctx *Context) JSON(statusCode int, data interface{}) {
	res := ctx.Res()
	res.SetStatusCode(statusCode)
	res.SetContentType("application/json")
	res.SetResult(ctx.View().JSON(data))
}

func (ctx *Context) HTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	res := ctx.Res()
	res.SetStatusCode(statusCode)
	res.SetContentType("text/html")
	for k, v := range renderArgs {
		ctx.SetRenderArg(k, v)
	}
	res.SetResult(ctx.View().HTML(tmplPath))
}

func (ctx *Context) Redirect(uri string) {
	res := ctx.Res()
	res.SetStatusCode(http.StatusFound)
	res.Header().Set("Location", uri)
	res.SetResult(ctx.View().Text(""))
}
