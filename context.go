package ksatriya

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Args struct {
	httprouter.Params
}
type Params map[string][]string

type Context struct {
	Request  *http.Request
	Response *Response
	View     *View
	Args     Args
	Params   Params
}

func NewContext(req *http.Request, args httprouter.Params) *Context {
	req.ParseForm()
	params := map[string][]string(req.Form)
	return &Context{
		Request:  req,
		Response: NewResponse(),
		View:     NewView(),
		Args:     Args{args},
		Params:   params,
	}
}

func (ctx *Context) SetTmplDirPath(val string) {
	ctx.View.RenderConfig.TmplDirPath = val
}

func (ctx *Context) SetBaseTmplPath(val string) {
	ctx.View.RenderConfig.BaseTmplPath = val
}

func (ctx *Context) SetRenderArg(key string, val interface{}) {
	ctx.View.RenderArgs[key] = val
}

func (ctx *Context) Arg(name string) string {
	return ctx.Args.ByName(name)
}

func (ctx *Context) Param(name string) ([]string, bool) {
	param, found := ctx.Params[name]
	return param, found
}

func (ctx *Context) ParamSingle(name string) (string, bool) {
	if param, found := ctx.Param(name); found {
		if len(param) > 0 {
			return param[0], true
		}
	}
	return "", false
}

func (ctx *Context) Text(stat int, text string) {
	res := ctx.Response
	res.StatusCode = stat
	res.SetContentType("text/plain")
	ctx.View.Renderer = NewTextRenderer(text)
}

func (ctx *Context) JSON(stat int, data interface{}) {
	res := ctx.Response
	res.StatusCode = stat
	res.SetContentType("application/json")
	ctx.View.Renderer = NewJSONRenderer(data)
}

func (ctx *Context) HTML(stat int, tmplPath string, args RenderArgs) {
	res := ctx.Response
	res.StatusCode = stat
	res.SetContentType("text/html")
	ctx.View.Renderer = NewHTMLRenderer(tmplPath, args)
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
