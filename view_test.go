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
	rText, err := tr.Render()
	assert.NoError(t, err)
	vText, err := v.Render()
	assert.NoError(t, err)
	assert.Equal(t, rText, vText)

	jr := NewJSONRenderer(map[string]string{"key": "value"})

	// set json renderer
	v.SetRenderer(jr)
	assert.Equal(t, jr, v.Renderer())

	// render json
	rJSON, err := jr.Render()
	assert.NoError(t, err)
	vJSON, err := v.Render()
	assert.NoError(t, err)
	assert.Equal(t, rJSON, vJSON)
}
