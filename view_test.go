package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	// new view
	v := NewView()
	assert.Zero(t, v.Renderer())

	// set text renderer
	tr := NewTextRenderer("text")
	v.SetRenderer(tr)
	assert.Equal(t, tr, v.Renderer())

	// render text
	assert.Equal(t, tr.Render(), v.Render())

	j := map[string]string{}
	j["key"] = "value"

	// set json renderer
	jr := NewJSONRenderer(j)
	v.SetRenderer(jr)
	assert.Equal(t, jr, v.Renderer())

	// render json
	assert.Equal(t, jr.Render(), v.Render())
}
