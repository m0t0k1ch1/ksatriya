package ksatriya

import "fmt"

const TmplDirDefault = "view"

type RenderData map[string]interface{}

type Renderer struct {
	TmplDir string
}

func NewRenderer() *Renderer {
	return &Renderer{
		TmplDir: TmplDirDefault,
	}
}

func (r *Renderer) RenderText(text string) Result {
	return &ResultText{
		Text: text,
	}
}

func (r *Renderer) RenderJSON(data interface{}) Result {
	return &ResultJSON{
		Data: data,
	}
}

func (r *Renderer) RenderHTML(tmplPath string, renderData *RenderData) Result {
	return &ResultHTML{
		TmplPath:   fmt.Sprintf("%s/%s", r.TmplDir, tmplPath),
		RenderData: renderData,
	}
}
