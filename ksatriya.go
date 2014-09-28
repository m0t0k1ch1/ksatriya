package ksatriya

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Params struct {
	httprouter.Params
}

type App interface {
	Router() *httprouter.Router
	View() ResultBuilder

	Handle(method, path string, handler HandlerFunc, filters map[string]FilterFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Run(addr string)
	RegisterController(d Dispacher)
}

type Ksatriya struct {
	router *httprouter.Router
	view   ResultBuilder
}

func New() *Ksatriya {
	k := &Ksatriya{
		router: httprouter.New(),
		view:   NewView(),
	}
	return k
}

func (k *Ksatriya) Router() *httprouter.Router {
	return k.router
}

func (k *Ksatriya) View() ResultBuilder {
	return k.view
}

func (k *Ksatriya) Handle(method, path string, handler HandlerFunc, filters map[string]FilterFunc) {
	k.Router().Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := NewContext(req, Params{params}, k.View())
		if filter, ok := filters[BeforeFilterKey]; ok {
			filter(ctx)
		}
		handler(ctx)
		if filter, ok := filters[AfterFilterKey]; ok {
			filter(ctx)
		}
		ctx.Res().Write(ctx, w)
	})
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.Router().ServeHTTP(w, req)
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.Router()); err != nil {
		panic(err)
	}
}

func (k *Ksatriya) RegisterController(d Dispacher) {
	for _, handler := range d.Routes() {
		k.Handle(handler.Method, handler.Path, handler.Func, d.Filters())
	}
}
