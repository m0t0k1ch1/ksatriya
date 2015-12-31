package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextRenderer(t *testing.T) {
	renderer := NewTextRenderer("text")

	assert.Equal(t, "text", renderer.Render())
}

func TestJSONRenderer(t *testing.T) {
	json := map[string]string{}
	json["key"] = "value"

	renderer := NewJSONRenderer(json)

	assert.Equal(t, "{\"key\":\"value\"}", renderer.Render())
}
