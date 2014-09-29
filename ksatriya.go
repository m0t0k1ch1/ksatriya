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

func (k *Ksatriya) Handle(method, path string, handler HandlerFunc, filters map[string]FilterFunc) {
	k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := NewContext(req, Params{params})

		if filter, ok := filters[BeforeFilterKey]; ok {
			filter(ctx)
			if ctx.Response.StatusCode == http.StatusFound {
				ctx.Write(w)
				return
			}
		}

		handler(ctx)
		if ctx.Response.StatusCode == http.StatusFound {
			ctx.Write(w)
			return
		}

		if filter, ok := filters[AfterFilterKey]; ok {
			filter(ctx)
		}

		ctx.Write(w)
	})
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.Router.ServeHTTP(w, req)
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.Router); err != nil {
		panic(err)
	}
}

func (k *Ksatriya) RegisterController(d Dispacher) {
	for _, handler := range d.Routes() {
		k.Handle(handler.Method, handler.Path, handler.Func, d.Filters())
	}
}
