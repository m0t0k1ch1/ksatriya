package ksatriya

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type Params struct {
	httprouter.Params
}

type Ksatriya struct {
	Router   *httprouter.Router
	Renderer *Renderer
}

func New() *Ksatriya {
	k := &Ksatriya{}
	k.Init()
	return k
}

func (k *Ksatriya) Init() {
	k.Router = httprouter.New()
	k.Renderer = NewRenderer()
}

func (k *Ksatriya) Run(addr string) {
	http.ListenAndServe(addr, k.Router)
}

func (k *Ksatriya) Handle(method, path string, handler HandlerFunc) {
	k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		c := NewContext(w, req, Params{params}, k.Renderer)
		handler(c)
	})
}

func (k *Ksatriya) GET(path string, handler HandlerFunc) {
	k.Handle("GET", path, handler)
}

func (k *Ksatriya) POST(path string, handler HandlerFunc) {
	k.Handle("POST", path, handler)
}

func (k *Ksatriya) PUT(path string, handler HandlerFunc) {
	k.Handle("PUT", path, handler)
}

func (k *Ksatriya) PATCH(path string, handler HandlerFunc) {
	k.Handle("PATCH", path, handler)
}

func (k *Ksatriya) DELETE(path string, handler HandlerFunc) {
	k.Handle("DELETE", path, handler)
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.Router.ServeHTTP(w, req)
}
