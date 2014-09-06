package ksatriya

import "fmt"

const TmplDirDefault = "view"

type RenderArgs map[string]interface{}

type Renderer struct {
	TmplDir      string
	BaseTmplPath string
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

func (r *Renderer) RenderHTML(tmplPath string) Result {
	result := &ResultHTML{
		TmplPath: fmt.Sprintf("%s/%s", r.TmplDir, tmplPath),
	}
	if len(r.BaseTmplPath) > 0 {
		result.BaseTmplPath = fmt.Sprintf("%s/%s", r.TmplDir, r.BaseTmplPath)
	}
	return result
}
