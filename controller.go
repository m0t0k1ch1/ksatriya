package ksatriya

const (
	BeforeFilterKey = "BEFORE"
	AfterFilterKey  = "AFTER"
)

type HandlerFunc func(*Context)
type FilterFunc func(*Context)

type Handler struct {
	Path   string
	Method string
	Func   HandlerFunc
}

type Dispacher interface {
	Routes() []*Handler
	Filters() map[string]FilterFunc

	AddRoute(method, path string, h HandlerFunc)
	GET(path string, h HandlerFunc)
	POST(path string, h HandlerFunc)
	PUT(path string, h HandlerFunc)
	PATCH(path string, h HandlerFunc)
	DELETE(path string, h HandlerFunc)

	AddBeforeFilter(f FilterFunc)
	AddAfterFilter(f FilterFunc)
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

func (c *Controller) AddRoute(method, path string, h HandlerFunc) {
	c.routes = append(c.routes, &Handler{
		Path:   path,
		Method: method,
		Func:   h,
	})
}

func (c *Controller) GET(path string, h HandlerFunc) {
	c.AddRoute("GET", path, h)
}

func (c *Controller) POST(path string, h HandlerFunc) {
	c.AddRoute("POST", path, h)
}

func (c *Controller) PUT(path string, h HandlerFunc) {
	c.AddRoute("PUT", path, h)
}

func (c *Controller) PATCH(path string, h HandlerFunc) {
	c.AddRoute("PATCH", path, h)
}

func (c *Controller) DELETE(path string, h HandlerFunc) {
	c.AddRoute("DELETE", path, h)
}

func (c *Controller) AddBeforeFilter(f FilterFunc) {
	c.filters[BeforeFilterKey] = f
}

func (c *Controller) AddAfterFilter(f FilterFunc) {
	c.filters[AfterFilterKey] = f
}
