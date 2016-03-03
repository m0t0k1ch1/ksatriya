package ksatriya

import "encoding/json"

type Renderer interface {
	Render() (string, error)
}

type TextRenderer struct {
	text string
}

func (r *TextRenderer) Render() (string, error) {
	return r.text, nil
}

func NewTextRenderer(text string) *TextRenderer {
	return &TextRenderer{
		text: text,
	}
}

type JSONRenderer struct {
	data interface{}
}

func (r *JSONRenderer) Render() (string, error) {
	b, err := json.Marshal(r.data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func NewJSONRenderer(data interface{}) *JSONRenderer {
	return &JSONRenderer{
		data: data,
	}
}
