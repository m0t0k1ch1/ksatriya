package ksatriya

import "encoding/json"

type Renderer interface {
	Render() string
}

type TextRenderer struct {
	text string
}

func NewTextRenderer(text string) *TextRenderer {
	return &TextRenderer{
		text: text,
	}
}

func (r *TextRenderer) Render() string {
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

func (r *JSONRenderer) Render() string {
	b, err := json.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	return string(b)
}
