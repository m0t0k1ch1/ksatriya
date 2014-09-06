package ksatriya

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

type Result interface {
	Apply(ctx Context, w http.ResponseWriter)
}

type ResultText struct {
	Text string
}

func (result *ResultText) Apply(ctx Context, w http.ResponseWriter) {
	w.Write([]byte(result.Text))
}

type ResultJSON struct {
	Data interface{}
}

func (result *ResultJSON) Apply(ctx Context, w http.ResponseWriter) {
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

func (result *ResultHTML) Apply(ctx Context, w http.ResponseWriter) {
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
