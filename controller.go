package ksatriya

type Dispacher interface {
	Routes() []*Handler

	AddRoute(method, path string, hf HandlerFunc)
	GET(path string, hf HandlerFunc)
	POST(path string, hf HandlerFunc)
	PUT(path string, hf HandlerFunc)
	PATCH(path string, hf HandlerFunc)
	DELETE(path string, hf HandlerFunc)
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
