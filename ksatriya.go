package ksatriya

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

type Params struct {
	httprouter.Params
}

type Ksatriya struct {
	Router *httprouter.Router
	View   *View
	DB     *gorm.DB
}

func New() *Ksatriya {
	k := &Ksatriya{}
	k.Init()
	return k
}

func (k *Ksatriya) Init() {
	k.Router = httprouter.New()
	k.View = NewView()
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.Router); err != nil {
		panic(err)
	}
}

func (k *Ksatriya) Handle(method, path string, handler HandlerFunc, filters map[string]FilterFunc) {
	k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ctx := NewContext(req, Params{params}, k.View, k.DB)
		if filter, ok := filters[BeforeFilterKey]; ok {
			filter(ctx)
		}
		handler(ctx)
		if filter, ok := filters[AfterFilterKey]; ok {
			filter(ctx)
		}
		ctx.Response.Write(ctx, w)
	})
}

func (k *Ksatriya) RegisterController(d Dispacher) {
	for _, handler := range d.Routes() {
		k.Handle(handler.Method, handler.Path, handler.Func, d.Filters())
	}
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.Router.ServeHTTP(w, req)
}
