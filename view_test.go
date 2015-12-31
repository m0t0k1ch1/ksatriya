package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	v := NewView()

	assert.Zero(t, v.Renderer())

	tr := NewTextRenderer("text")
	v.SetRenderer(tr)

	assert.Equal(t, tr.Render(), v.Render())

	json := map[string]string{}
	json["key"] = "value"
	jr := NewJSONRenderer(json)
	v.SetRenderer(jr)

	assert.Equal(t, jr.Render(), v.Render())
}
