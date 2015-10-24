package ksatriya

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	RootPathDefault = "/static"
	RootDirDefault  = "static"
)

type Root struct {
	path string
	dir  http.FileSystem
}

func NewRoot() *Root {
	root := &Root{}
	root.SetPath(RootPathDefault)
	root.SetDir(RootDirDefault)
	return root
}

func (root *Root) Path() string {
	return root.path
}
func (root *Root) SetPath(val string) {
	root.path = val + "/*filepath"
}

func (root *Root) Dir() http.FileSystem {
	return root.dir
}
func (root *Root) SetDir(val string) {
	root.dir = http.Dir(val)
}

type Ksatriya struct {
	router *httprouter.Router
	root   *Root
}

func New() *Ksatriya {
	return &Ksatriya{
		router: httprouter.New(),
		root:   NewRoot(),
	}
}

func (k *Ksatriya) Root() *Root {
	return k.root
}

func (k *Ksatriya) handle(method, path string, hf HandlerFunc, filterFuncs map[string]FilterFunc) {
	k.router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, args httprouter.Params) {
		ctx := NewContext(req, args)

		if ff, ok := filterFuncs[BeforeFilterFuncKey]; ok {
			ff(ctx)
			if ctx.Res().StatusCode() == http.StatusFound {
				ctx.Write(w)
				return
			}
		}

		hf(ctx)
		if ctx.Res().StatusCode() == http.StatusFound {
			ctx.Write(w)
			return
		}

		if ff, ok := filterFuncs[AfterFilterFuncKey]; ok {
			ff(ctx)
		}

		ctx.Write(w)
	})
}

func (k *Ksatriya) AddRoute(method, path string, hf HandlerFunc) {
	k.handle(method, path, hf, map[string]FilterFunc{})
}

func (k *Ksatriya) GET(path string, hf HandlerFunc) {
	k.AddRoute("GET", path, hf)
}

func (k *Ksatriya) POST(path string, hf HandlerFunc) {
	k.AddRoute("POST", path, hf)
}

func (k *Ksatriya) PUT(path string, hf HandlerFunc) {
	k.AddRoute("PUT", path, hf)
}

func (k *Ksatriya) PATCH(path string, hf HandlerFunc) {
	k.AddRoute("PATCH", path, hf)
}

func (k *Ksatriya) DELETE(path string, hf HandlerFunc) {
	k.AddRoute("DELETE", path, hf)
}

func (k *Ksatriya) RegisterController(d Dispacher) {
	for _, h := range d.Routes() {
		k.handle(h.Method(), h.Path(), h.HandlerFunc(), d.FilterFuncs())
	}
}

func (k *Ksatriya) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	k.router.ServeHTTP(w, req)
}

func (k *Ksatriya) ServeFiles() {
	root := k.Root()
	k.router.ServeFiles(root.Path(), root.Dir())
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.router); err != nil {
		panic(err)
	}
}
