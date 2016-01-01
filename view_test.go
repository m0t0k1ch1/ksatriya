package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	// new view
	v := NewView()
	assert.Empty(t, v.Renderer())

	tr := NewTextRenderer("text")

	// set text renderer
	v.SetRenderer(tr)
	assert.Equal(t, tr, v.Renderer())

	// render text
	assert.Equal(t, tr.Render(), v.Render())

	jr := NewJSONRenderer(map[string]string{"key": "value"})

	// set json renderer
	v.SetRenderer(jr)
	assert.Equal(t, jr, v.Renderer())

	// render json
	assert.Equal(t, jr.Render(), v.Render())
}
