package ksatriya

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	RootPathDefault = "/static"
	RootDirDefault  = "static"
)

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}

type Root struct {
	path string
	dir  http.FileSystem
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

func NewRoot() *Root {
	root := &Root{}
	root.SetPath(RootPathDefault)
	root.SetDir(RootDirDefault)
	return root
}

type Ksatriya struct {
	router     *Router
	root       *Root
	ctxBuilder CtxBuilder
}

func (k *Ksatriya) Router() *Router {
	return k.router
}

func (k *Ksatriya) Root() *Root {
	return k.root
}

func (k *Ksatriya) SetCtxBuilder(f CtxBuilder) {
	k.ctxBuilder = f
}

func (k *Ksatriya) handle(method, path string, hf HandlerFunc, filterFuncs map[string]FilterFunc) {
	k.Router().Handle(method, path, func(w http.ResponseWriter, req *http.Request, args httprouter.Params) {
		ctx := k.ctxBuilder(w, req, Args{args})
		defer ctx.Finalize()

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
	k.Router().ServeHTTP(w, req)
}

func (k *Ksatriya) ServeFiles() {
	root := k.Root()
	k.Router().ServeFiles(root.Path(), root.Dir())
}

func (k *Ksatriya) Run(addr string) {
	if err := http.ListenAndServe(addr, k.Router()); err != nil {
		log.Fatal(err)
	}
}

func New() *Ksatriya {
	return &Ksatriya{
		router:     NewRouter(),
		root:       NewRoot(),
		ctxBuilder: NewContext,
	}
}
