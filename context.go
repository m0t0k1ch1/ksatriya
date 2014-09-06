package ksatriya

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Context struct {
	Request    *http.Request
	Response   *Response
	Params     Params
	Renderer   *Renderer
	RenderArgs RenderArgs
	DB         gorm.DB
}

func NewContext(req *http.Request, params Params, r *Renderer, db gorm.DB) *Context {
	req.ParseForm()
	return &Context{
		Request: req,
		Response: &Response{
			StatusCode: 200,
			Header:     http.Header{},
		},
		Params:     params,
		Renderer:   r,
		RenderArgs: RenderArgs{},
		DB:         db,
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

func (ctx *Context) SetBaseTmplPath(baseTmplPath string) {
	ctx.Renderer.BaseTmplPath = baseTmplPath
}

func (ctx *Context) RenderText(statusCode int, text string) {
	ctx.SetStatusCode(statusCode)
	ctx.Response.Result = ctx.Renderer.RenderText(text)
}

func (ctx *Context) RenderJSON(statusCode int, data interface{}) {
	ctx.SetStatusCode(statusCode)
	ctx.Response.Result = ctx.Renderer.RenderJSON(data)
}

func (ctx *Context) RenderHTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	ctx.SetStatusCode(statusCode)
	for k, v := range renderArgs {
		ctx.RenderArgs[k] = v
	}
	ctx.Response.Result = ctx.Renderer.RenderHTML(tmplPath)
}

func (ctx *Context) Write(w http.ResponseWriter) {
	ctx.Response.Write(ctx, w)
}
