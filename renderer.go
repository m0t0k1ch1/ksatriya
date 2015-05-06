package ksatriya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"text/template"
)

const (
	TmplDirPathDefault  = "view"
	BaseTmplPathDefault = ""
)

type RenderArgs map[string]interface{}

type RenderConfig struct {
	tmplDirPath  string
	baseTmplPath string
}

func NewRenderConfig() *RenderConfig {
	return &RenderConfig{
		tmplDirPath:  TmplDirPathDefault,
		baseTmplPath: BaseTmplPathDefault,
	}
}

func (conf *RenderConfig) TmplDirPath() string {
	return conf.tmplDirPath
}
func (conf *RenderConfig) SetTmplDirPath(val string) {
	conf.tmplDirPath = val
}

func (conf *RenderConfig) BaseTmplPath() string {
	return conf.baseTmplPath
}
func (conf *RenderConfig) SetBaseTmplPath(val string) {
	conf.baseTmplPath = val
}

type Renderer interface {
	Render(conf *RenderConfig, args RenderArgs) string
}

type TextRenderer struct {
	text string
}

func NewTextRenderer(text string) *TextRenderer {
	return &TextRenderer{
		text: text,
	}
}

func (r *TextRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	return r.text
}

type JSONRenderer struct {
	data interface{}
}

func NewJSONRenderer(data interface{}) *JSONRenderer {
	return &JSONRenderer{
		data: data,
	}
}

func (r *JSONRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	b, err := json.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

type HTMLRenderer struct {
	tmplPath   string
	renderArgs RenderArgs
}

func NewHTMLRenderer(tmplPath string, args RenderArgs) *HTMLRenderer {
	return &HTMLRenderer{
		tmplPath:   tmplPath,
		renderArgs: args,
	}
}

func (r *HTMLRenderer) Render(conf *RenderConfig, args RenderArgs) string {
	for key, val := range r.renderArgs {
		args[key] = val
	}

	tmplPath := fmt.Sprintf("%s/%s", conf.TmplDirPath(), r.tmplPath)

	var tmpl *template.Template
	if len(conf.BaseTmplPath()) > 0 {
		baseTmplPath := fmt.Sprintf("%s/%s", conf.TmplDirPath(), conf.BaseTmplPath())
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
