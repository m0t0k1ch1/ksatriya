package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	v := NewView()

	assert.Zero(t, v.Renderer())

	textRenderer := NewTextRenderer("text")
	v.SetRenderer(textRenderer)

	assert.Equal(t, textRenderer.Render(), v.Render())

	json := map[string]string{}
	json["key"] = "value"
	jsonRenderer := NewJSONRenderer(json)
	v.SetRenderer(jsonRenderer)

	assert.Equal(t, jsonRenderer.Render(), v.Render())
}
