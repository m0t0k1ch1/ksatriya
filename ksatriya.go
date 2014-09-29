package ksatriya

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Params struct {
	httprouter.Params
}

type Ksatriya struct {
	Router *httprouter.Router
}

func New() *Ksatriya {
	return &Ksatriya{
		Router: httprouter.New(),
	}
}

func (k *Ksatriya) Handle(method, path string, h HandlerFunc, filters map[string]FilterFunc) {
	k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := NewContext(req, Params{params})

		if f, ok := filters[BeforeFilterKey]; ok {
			f(ctx)
			if ctx.Response.StatusCode == http.StatusFound {
				ctx.Write(w)
				return
			}
		}

		h(ctx)
		if ctx.Response.StatusCode == http.StatusFound {
			ctx.Write(w)
			return
		}

		if f, ok := filters[AfterFilterKey]; ok {
			f(ctx)
		}

		ctx.Write(w)
	})
}

func (k *Ksatriya) AddRoute(method, path string, h HandlerFunc) {
	k.Handle(method, path, h, map[string]FilterFunc{})
}

func (k *Ksatriya) GET(path string, h HandlerFunc) {
	k.AddRoute("GET", path, h)
}

func (k *Ksatriya) POST(path string, h HandlerFunc) {
	k.AddRoute("POST", path, h)
}

func (k *Ksatriya) PUT(path string, h HandlerFunc) {
	k.AddRoute("PUT", path, h)
}

func (k *Ksatriya) PATCH(path string, h HandlerFunc) {
	k.AddRoute("PATCH", path, h)
}

func (k *Ksatriya) DELETE(path string, h HandlerFunc) {
	k.AddRoute("DELETE", path, h)
}

func (k *Ksatriya) RegisterController(d Dispacher) {
	for _, handler := range d.Routes() {
		k.Handle(handler.Method, handler.Path, handler.Func, d.Filters())
	}
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.Router.ServeHTTP(w, req)
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.Router); err != nil {
		panic(err)
	}
}
