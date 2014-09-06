package ksatriya

const (
	BeforeFilterKey = "BEFORE"
	AfterFilterKey  = "AFTER"
)

type HandlerFunc func(Context)
type FilterFunc func(Context)

type Handler struct {
	Path   string
	Method string
	Func   HandlerFunc
}

type Dispacher interface {
	Routes() []*Handler
	Filters() map[string]FilterFunc
	GET(path string, handlerFunc HandlerFunc)
	POST(path string, handlerFunc HandlerFunc)
	PUT(path string, handlerFunc HandlerFunc)
	PATCH(path string, handlerFunc HandlerFunc)
	DELETE(path string, handlerFunc HandlerFunc)
	AddRoute(method, path string, handlerFunc HandlerFunc)
	AddBeforeFilter(filterFunc FilterFunc)
	AddAfterFilter(filterFunc FilterFunc)
}

type Controller struct {
	routes  []*Handler
	filters map[string]FilterFunc
}

func NewController() *Controller {
	return &Controller{
		routes:  []*Handler{},
		filters: map[string]FilterFunc{},
	}
}

func (c *Controller) Routes() []*Handler {
	return c.routes
}

func (c *Controller) Filters() map[string]FilterFunc {
	return c.filters
}

func (c *Controller) GET(path string, handlerFunc HandlerFunc) {
	c.AddRoute("GET", path, handlerFunc)
}

func (c *Controller) POST(path string, handlerFunc HandlerFunc) {
	c.AddRoute("POST", path, handlerFunc)
}

func (c *Controller) PUT(path string, handlerFunc HandlerFunc) {
	c.AddRoute("PUT", path, handlerFunc)
}

func (c *Controller) PATCH(path string, handlerFunc HandlerFunc) {
	c.AddRoute("PATCH", path, handlerFunc)
}

func (c *Controller) DELETE(path string, handlerFunc HandlerFunc) {
	c.AddRoute("DELETE", path, handlerFunc)
}

func (c *Controller) AddRoute(method, path string, handlerFunc HandlerFunc) {
	c.routes = append(c.routes, &Handler{
		Path:   path,
		Method: method,
		Func:   handlerFunc,
	})
}

func (c *Controller) AddBeforeFilter(filterFunc FilterFunc) {
	c.filters[BeforeFilterKey] = filterFunc
}

func (c *Controller) AddAfterFilter(filterFunc FilterFunc) {
	c.filters[AfterFilterKey] = filterFunc
}
