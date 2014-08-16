package ksatriya

import (
	"net/http"

	"github.com/unrolled/render"
)

const LayoutDefault = "layout"

type RenderData map[string]interface{}

type Renderer struct {
	*render.Render
}

func NewRenderer() *Renderer {
	return &Renderer{render.New(render.Options{})}
}

func (r *Renderer) RenderText(w http.ResponseWriter, status int, text string) {
	w.Write([]byte(text))
}

func (r *Renderer) RenderHTML(w http.ResponseWriter, status int, name string, data RenderData, layout []string) {
	htmlOptions := render.HTMLOptions{Layout: LayoutDefault}
	if len(layout) > 0 {
		htmlOptions.Layout = layout[0]
	}
	r.HTML(w, status, name, data, htmlOptions)
}

func (r *Renderer) RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	r.JSON(w, status, v)
}
