package ksatriya

const (
	BeforeFilterFuncKey = "BEFORE"
	AfterFilterFuncKey  = "AFTER"
)

type HandlerFunc func(Ctx)
type FilterFunc func(Ctx)

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

type Dispacher interface {
	Routes() []*Handler
	FilterFuncs() map[string]FilterFunc

	AddRoute(method, path string, hf HandlerFunc)
	GET(path string, hf HandlerFunc)
	POST(path string, hf HandlerFunc)
	PUT(path string, hf HandlerFunc)
	PATCH(path string, hf HandlerFunc)
	DELETE(path string, hf HandlerFunc)

	AddBeforeFilter(ff FilterFunc)
	AddAfterFilter(ff FilterFunc)
}

type Controller struct {
	routes      []*Handler
	filterFuncs map[string]FilterFunc
}

func NewController() *Controller {
	return &Controller{
		routes:      []*Handler{},
		filterFuncs: map[string]FilterFunc{},
	}
}

func (c *Controller) Routes() []*Handler {
	return c.routes
}

func (c *Controller) FilterFuncs() map[string]FilterFunc {
	return c.filterFuncs
}

func (c *Controller) AddRoute(method, path string, hf HandlerFunc) {
	h := &Handler{
		method:      method,
		path:        path,
		handlerFunc: hf,
	}
	c.routes = append(c.routes, h)
}

func (c *Controller) GET(path string, hf HandlerFunc) {
	c.AddRoute("GET", path, hf)
}

func (c *Controller) POST(path string, hf HandlerFunc) {
	c.AddRoute("POST", path, hf)
}

func (c *Controller) PUT(path string, hf HandlerFunc) {
	c.AddRoute("PUT", path, hf)
}

func (c *Controller) PATCH(path string, hf HandlerFunc) {
	c.AddRoute("PATCH", path, hf)
}

func (c *Controller) DELETE(path string, hf HandlerFunc) {
	c.AddRoute("DELETE", path, hf)
}

func (c *Controller) AddBeforeFilter(ff FilterFunc) {
	c.filterFuncs[BeforeFilterFuncKey] = ff
}

func (c *Controller) AddAfterFilter(ff FilterFunc) {
	c.filterFuncs[AfterFilterFuncKey] = ff
}
