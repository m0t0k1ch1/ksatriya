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
	request  *http.Request
	response *Response
	view     *View
	args     Args
	params   Params
}

func NewContext(req *http.Request, args httprouter.Params) *Context {
	req.ParseForm()
	params := map[string][]string(req.Form)

	return &Context{
		request:  req,
		response: NewResponse(),
		view:     NewView(),
		args:     Args{args},
		params:   params,
	}
}

func (ctx *Context) Req() *http.Request {
	return ctx.request
}

func (ctx *Context) Res() *Response {
	return ctx.response
}

func (ctx *Context) View() *View {
	return ctx.view
}

func (ctx *Context) Args() Args {
	return ctx.args
}

func (ctx *Context) Arg(name string) string {
	return ctx.Args().ByName(name)
}

func (ctx *Context) Params() Params {
	return ctx.params
}

func (ctx *Context) Param(name string) ([]string, bool) {
	params := ctx.Params()
	param, found := params[name]
	return param, found
}

func (ctx *Context) ParamSingle(name string) string {
	if param, found := ctx.Param(name); found {
		if len(param) > 0 {
			return param[0]
		}
	}
	return ""
}

func (ctx *Context) Text(stat int, text string) {
	res := ctx.Res()
	res.SetStatusCode(stat)
	res.SetContentType("text/plain")
	ctx.View().SetRenderer(NewTextRenderer(text))
}

func (ctx *Context) JSON(stat int, data interface{}) {
	res := ctx.Res()
	res.SetStatusCode(stat)
	res.SetContentType("application/json")
	ctx.View().SetRenderer(NewJSONRenderer(data))
}

func (ctx *Context) HTML(stat int, tmplPath string, args RenderArgs) {
	res := ctx.Res()
	res.SetStatusCode(stat)
	res.SetContentType("text/html")
	ctx.View().SetRenderer(NewHTMLRenderer(tmplPath, args))
}

func (ctx *Context) Redirect(uri string) {
	res := ctx.Res()
	res.SetStatusCode(http.StatusFound)
	res.SetHeader("Location", uri)
	ctx.View().SetRenderer(NewTextRenderer(""))
}

func (ctx *Context) Write(w http.ResponseWriter) {
	res := ctx.Res()
	res.SetBody(ctx.View().Render())
	res.Write(w)
}
