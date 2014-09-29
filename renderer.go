package ksatriya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"text/template"
)

type Renderer interface {
	Render(conf *RenderConfig, args RenderArgs) string
}

type TextRenderer struct {
	Text string
}

func NewTextRenderer(text string) *TextRenderer {
	return &TextRenderer{
		Text: text,
	}
}

func (r *TextRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	return r.Text
}

type JSONRenderer struct {
	Data interface{}
}

func NewJSONRenderer(data interface{}) *JSONRenderer {
	return &JSONRenderer{
		Data: data,
	}
}

func (r *JSONRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	b, err := json.Marshal(r.Data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

type HTMLRenderer struct {
	TmplPath   string
	RenderArgs RenderArgs
}

func NewHTMLRenderer(tmplPath string, args RenderArgs) *HTMLRenderer {
	return &HTMLRenderer{
		TmplPath:   tmplPath,
		RenderArgs: args,
	}
}

func (r *HTMLRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	for key, val := range r.RenderArgs {
		args[key] = val
	}

	tmplPath := fmt.Sprintf("%s/%s", conf.TmplDirPath, r.TmplPath)

	var tmpl *template.Template
	if len(conf.BaseTmplPath) > 0 {
		baseTmplPath := fmt.Sprintf("%s/%s", conf.TmplDirPath, conf.BaseTmplPath)
		tmpl = template.Must(template.New(filepath.Base(baseTmplPath)).ParseFiles(baseTmplPath, tmplPath))
	} else {
		tmpl = template.Must(template.New(filepath.Base(tmplPath)).ParseFiles(tmplPath))
	}

	b := &bytes.Buffer{}
	err := tmpl.Execute(b, args)
	if err != nil {
		panic(err)
	}

	return b.String()
}
