package ksatriya

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

type Result interface {
	Apply(ctx Ctx, w http.ResponseWriter)
}

type ResultText struct {
	Text string
}

func NewResultText(text string) *ResultText {
	return &ResultText{
		Text: text,
	}
}

func (result *ResultText) Apply(ctx Ctx, w http.ResponseWriter) {
	w.Write([]byte(result.Text))
}

type ResultJSON struct {
	Data interface{}
}

func NewResultJSON(data interface{}) *ResultJSON {
	return &ResultJSON{
		Data: data,
	}
}

func (result *ResultJSON) Apply(ctx Ctx, w http.ResponseWriter) {
	b, err := json.Marshal(result.Data)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

type ResultHTML struct {
	BaseTmplPath string
	TmplPath     string
}

func NewResultHTML(tmplPath string) *ResultHTML {
	return &ResultHTML{
		TmplPath: tmplPath,
	}
}

func (result *ResultHTML) Apply(ctx Ctx, w http.ResponseWriter) {
	buffer := &bytes.Buffer{}
	var tmpl *template.Template
	if len(result.BaseTmplPath) > 0 {
		tmpl = template.Must(template.New(filepath.Base(result.BaseTmplPath)).ParseFiles(result.BaseTmplPath, result.TmplPath))
	} else {
		tmpl = template.Must(template.New(filepath.Base(result.TmplPath)).ParseFiles(result.TmplPath))
	}
	err := tmpl.Execute(w, ctx.RenderArgs())
	if err != nil {
		panic(err)
	}
	w.Write(buffer.Bytes())
}
