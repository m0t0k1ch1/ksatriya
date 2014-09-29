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
	View   ResultBuilder
}

func New() *Ksatriya {
	k := &Ksatriya{
		Router: httprouter.New(),
		View:   NewView(),
	}
	return k
}

func (k *Ksatriya) Handle(method, path string, handler HandlerFunc, filters map[string]FilterFunc) {
	k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := NewContext(w, req, Params{params}, k.View)

		if filter, ok := filters[BeforeFilterKey]; ok {
			filter(ctx)
			if ctx.Res().StatusCode() == http.StatusFound {
				ctx.Res().Write(ctx)
				return
			}
		}

		handler(ctx)
		if ctx.Res().StatusCode() == http.StatusFound {
			ctx.Res().Write(ctx)
			return
		}

		if filter, ok := filters[AfterFilterKey]; ok {
			filter(ctx)
		}

		ctx.Res().Write(ctx)
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
