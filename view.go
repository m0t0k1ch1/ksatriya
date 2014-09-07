package ksatriya

import "fmt"

const TmplDirPathDefault = "view"

type RenderArgs map[string]interface{}

type View struct {
	TmplDirPath  string
	BaseTmplPath string
}

func NewView() *View {
	return &View{
		TmplDirPath: TmplDirPathDefault,
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
		TmplPath: fmt.Sprintf("%s/%s", v.TmplDirPath, tmplPath),
	}
	if len(v.BaseTmplPath) > 0 {
		result.BaseTmplPath = fmt.Sprintf("%s/%s", v.TmplDirPath, v.BaseTmplPath)
	}
	return result
}
