package ksatriya

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(Ctx)

type Handler struct {
	method      string
	path        string
	handlerFunc HandlerFunc
}

func (h *Handler) Method() string {
	return h.method
}

func (h *Handler) Path() string {
	return h.path
}

func (h *Handler) HandlerFunc() HandlerFunc {
	return h.handlerFunc
}

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
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

func (k *Ksatriya) AddRoute(method, path string, hf HandlerFunc) {
	k.Router().Handle(method, path, func(w http.ResponseWriter, req *http.Request, args httprouter.Params) {
		ctx := k.ctxBuilder(w, req, Args{args})
		defer ctx.Finalize()

		hf(ctx)
		if ctx.Res().StatusCode() == http.StatusFound {
			ctx.Write(w)
			return
		}

		ctx.Write(w)
	})
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
		k.AddRoute(h.Method(), h.Path(), h.HandlerFunc())
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
