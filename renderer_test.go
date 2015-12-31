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
	j := map[string]string{}
	j["key"] = "value"
	r := NewJSONRenderer(j)

	assert.Equal(t, "{\"key\":\"value\"}", r.Render())
}
