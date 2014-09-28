package ksatriya

import "fmt"

const TmplDirPathDefault = "view"

type RenderArgs map[string]interface{}

type ResultBuilder interface {
	TmplDirPath() string
	BaseTmplPath() string

	SetTmplDirPath(path string)
	SetBaseTmplPath(path string)

	Text(text string) Result
	JSON(data interface{}) Result
	HTML(tmplPath string) Result
}

type View struct {
	tmplDirPath  string
	baseTmplPath string
}

func NewView() *View {
	return &View{
		tmplDirPath: TmplDirPathDefault,
	}
}

func (v *View) TmplDirPath() string {
	return v.tmplDirPath
}

func (v *View) BaseTmplPath() string {
	return v.baseTmplPath
}

func (v *View) SetTmplDirPath(path string) {
	v.tmplDirPath = path
}

func (v *View) SetBaseTmplPath(path string) {
	v.baseTmplPath = path
}

func (v *View) Text(text string) Result {
	return NewResultText(text)
}

func (v *View) JSON(data interface{}) Result {
	return NewResultJSON(data)
}

func (v *View) HTML(tmplPath string) Result {
	result := NewResultHTML(fmt.Sprintf("%s/%s", v.TmplDirPath(), tmplPath))
	if len(v.BaseTmplPath()) > 0 {
		result.BaseTmplPath = fmt.Sprintf("%s/%s", v.TmplDirPath(), v.BaseTmplPath())
	}
	return result
}
