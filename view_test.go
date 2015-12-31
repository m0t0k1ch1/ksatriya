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

	j := map[string]string{}
	j["key"] = "value"
	jr := NewJSONRenderer(j)
	v.SetRenderer(jr)

	assert.Equal(t, jr.Render(), v.Render())
}
