package ksatriya

import "fmt"

const TmplDirDefault = "view"

type RenderArgs map[string]interface{}

type View struct {
	TmplDir      string
	BaseTmplPath string
}

func NewView() *View {
	return &View{
		TmplDir: TmplDirDefault,
	}
}

func (v *View) Text(text string) Result {
	return &ResultText{
		Text: text,
	}
}

func (v *View) JSON(data interface{}) Result {
	return &ResultJSON{
		Data: data,
	}
}

func (v *View) HTML(tmplPath string) Result {
	result := &ResultHTML{
		TmplPath: fmt.Sprintf("%s/%s", v.TmplDir, tmplPath),
	}
	if len(v.BaseTmplPath) > 0 {
		result.BaseTmplPath = fmt.Sprintf("%s/%s", v.TmplDir, v.BaseTmplPath)
	}
	return result
}
