package ksatriya

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
}

type Controller struct {
	routes []*Handler
}

func (c *Controller) Routes() []*Handler {
	return c.routes
}

func (c *Controller) AddRoute(method, path string, hf HandlerFunc) {
	h := &Handler{
		method:      method,
		path:        path,
		handlerFunc: hf,
	}
	c.routes = append(c.Routes(), h)
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

func NewController() *Controller {
	return &Controller{
		routes: []*Handler{},
	}
}
