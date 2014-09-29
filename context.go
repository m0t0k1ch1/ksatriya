package ksatriya

import "net/http"

type Context struct {
	Request  *http.Request
	Response *Response
	Params   Params
	View     *View
}

func NewContext(req *http.Request, params Params) *Context {
	req.ParseForm()
	return &Context{
		Request:  req,
		Response: NewResponse(),
		Params:   params,
		View:     NewView(),
	}
}

func (ctx *Context) SetTmplDirPath(path string) {
	ctx.View.RenderConfig.TmplDirPath = path
}

func (ctx *Context) SetBaseTmplPath(path string) {
	ctx.View.RenderConfig.BaseTmplPath = path
}

func (ctx *Context) SetRenderArg(key string, val interface{}) {
	ctx.View.RenderArgs[key] = val
}

func (ctx *Context) Param(name string) string {
	return ctx.Params.ByName(name)
}

func (ctx *Context) Text(statusCode int, text string) {
	res := ctx.Response
	res.StatusCode = statusCode
	res.SetContentType("text/plain")
	ctx.View.Renderer = NewTextRenderer(text)
}

func (ctx *Context) JSON(statusCode int, data interface{}) {
	res := ctx.Response
	res.StatusCode = statusCode
	res.SetContentType("application/json")
	ctx.View.Renderer = NewJSONRenderer(data)
}

func (ctx *Context) HTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	res := ctx.Response
	res.StatusCode = statusCode
	res.SetContentType("text/html")
	ctx.View.Renderer = NewHTMLRenderer(tmplPath, renderArgs)
}

func (ctx *Context) Redirect(uri string) {
	res := ctx.Response
	res.StatusCode = http.StatusFound
	res.Header.Set("Location", uri)
	ctx.View.Renderer = NewTextRenderer("")
}

func (ctx *Context) Write(w http.ResponseWriter) {
	res := ctx.Response
	res.Body = ctx.View.Render()
	res.Write(w)
}
