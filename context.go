package ksatriya

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Ctx interface {
	Req() *http.Request
	Res() *Response
	Params() Params
	Param(name string) string
	View() *View
	RenderArgs() RenderArgs

	SetStatusCode(statusCode int)
	SetRenderArg(key string, value interface{})
	SetTmplDirPath(tmplDirPath string)
	SetBaseTmplPath(baseTmplPath string)

	Text(statusCode int, text string)
	JSON(statusCode int, data interface{})
	HTML(statusCode int, tmplPath string, renderArgs RenderArgs)
	Redirect(uri string)
	Write(w http.ResponseWriter)
}

type Context struct {
	request    *http.Request
	response   *Response
	params     Params
	view       *View
	renderArgs RenderArgs
}

func NewContext(req *http.Request, params Params, v *View, db *gorm.DB) *Context {
	req.ParseForm()
	return &Context{
		request: req,
		response: &Response{
			StatusCode: 200,
			Header:     http.Header{},
		},
		params:     params,
		view:       v,
		renderArgs: RenderArgs{},
	}
}

func (ctx *Context) Req() *http.Request {
	return ctx.request
}

func (ctx *Context) Res() *Response {
	return ctx.response
}

func (ctx *Context) Params() Params {
	return ctx.params
}

func (ctx *Context) Param(name string) string {
	return ctx.Params().ByName(name)
}

func (ctx *Context) View() *View {
	return ctx.view
}

func (ctx *Context) RenderArgs() RenderArgs {
	return ctx.renderArgs
}

func (ctx *Context) SetStatusCode(statusCode int) {
	ctx.Res().StatusCode = statusCode
}

func (ctx *Context) SetRenderArg(key string, value interface{}) {
	ctx.renderArgs[key] = value
}

func (ctx *Context) SetTmplDirPath(tmplDirPath string) {
	ctx.View().TmplDirPath = tmplDirPath
}

func (ctx *Context) SetBaseTmplPath(baseTmplPath string) {
	ctx.View().BaseTmplPath = baseTmplPath
}

func (ctx *Context) Text(statusCode int, text string) {
	ctx.SetStatusCode(statusCode)
	ctx.Res().Result = ctx.View().Text(text)
}

func (ctx *Context) JSON(statusCode int, data interface{}) {
	ctx.SetStatusCode(statusCode)
	ctx.Res().Result = ctx.View().JSON(data)
}

func (ctx *Context) HTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	ctx.SetStatusCode(statusCode)
	for k, v := range renderArgs {
		ctx.SetRenderArg(k, v)
	}
	ctx.Res().Result = ctx.View().HTML(tmplPath)
}

func (ctx *Context) Redirect(uri string) {
	ctx.SetStatusCode(http.StatusFound)
	ctx.Res().Header.Set("Location", uri)
	ctx.Res().Result = ctx.View().Text("")
}

func (ctx *Context) Write(w http.ResponseWriter) {
	ctx.Res().Write(ctx, w)
}
