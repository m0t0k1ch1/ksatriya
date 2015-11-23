package ksatriya

import "encoding/json"

type Renderer interface {
	Render() string
}

type TextRenderer struct {
	text string
}

func (r *TextRenderer) Render() string {
	return r.text
}

func NewTextRenderer(text string) *TextRenderer {
	return &TextRenderer{
		text: text,
	}
}

type JSONRenderer struct {
	data interface{}
}

func (r *JSONRenderer) Render() string {
	b, err := json.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func NewJSONRenderer(data interface{}) *JSONRenderer {
	return &JSONRenderer{
		data: data,
	}
}
