package ksatriya

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

type Result interface {
	Apply(w http.ResponseWriter)
}

type ResultText struct {
	Text string
}

func (result *ResultText) Apply(w http.ResponseWriter) {
	w.Write([]byte(result.Text))
}

type ResultJSON struct {
	Data interface{}
}

func (result *ResultJSON) Apply(w http.ResponseWriter) {
	b, err := json.Marshal(result.Data)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

type ResultHTML struct {
	TmplPath   string
	RenderData *RenderData
}

func (result *ResultHTML) Apply(w http.ResponseWriter) {
	buffer := &bytes.Buffer{}
	tmpl := template.Must(template.New(filepath.Base(result.TmplPath)).ParseFiles(result.TmplPath))
	err := tmpl.Execute(w, result.RenderData)
	if err != nil {
		panic(err)
	}
	w.Write(buffer.Bytes())
}
