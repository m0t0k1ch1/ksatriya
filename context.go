package ksatriya

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Context struct {
	Request    *http.Request
	Response   *Response
	Params     Params
	View       *View
	RenderArgs RenderArgs
	DB         *gorm.DB
}

func NewContext(req *http.Request, params Params, v *View, db *gorm.DB) *Context {
	req.ParseForm()
	return &Context{
		Request: req,
		Response: &Response{
			StatusCode: 200,
			Header:     http.Header{},
		},
		Params:     params,
		View:       v,
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
	ctx.View.TmplDir = tmplDir
}

func (ctx *Context) SetBaseTmplPath(baseTmplPath string) {
	ctx.View.BaseTmplPath = baseTmplPath
}

func (ctx *Context) Text(statusCode int, text string) {
	ctx.SetStatusCode(statusCode)
	ctx.Response.Result = ctx.View.Text(text)
}

func (ctx *Context) JSON(statusCode int, data interface{}) {
	ctx.SetStatusCode(statusCode)
	ctx.Response.Result = ctx.View.JSON(data)
}

func (ctx *Context) HTML(statusCode int, tmplPath string, renderArgs RenderArgs) {
	ctx.SetStatusCode(statusCode)
	for k, v := range renderArgs {
		ctx.RenderArgs[k] = v
	}
	ctx.Response.Result = ctx.View.HTML(tmplPath)
}

func (ctx *Context) Write(w http.ResponseWriter) {
	ctx.Response.Write(ctx, w)
}
