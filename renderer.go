package ksatriya

const TmplPathBase = "template/"

type RenderData map[string]interface{}

type Renderer struct {
	BaseTmplPath string
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) SetBaseTmpl(baseTmplPath string) {
	r.BaseTmplPath = baseTmplPath
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
		TmplPath:   TmplPathBase + tmplPath,
		RenderData: renderData,
	}
}
