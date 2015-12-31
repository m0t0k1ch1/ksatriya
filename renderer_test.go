package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextRenderer(t *testing.T) {
	r := NewTextRenderer("text")

	assert.Equal(t, "text", r.Render())
}

func TestJSONRenderer(t *testing.T) {
	json := map[string]string{}
	json["key"] = "value"
	r := NewJSONRenderer(json)

	assert.Equal(t, "{\"key\":\"value\"}", r.Render())
}
