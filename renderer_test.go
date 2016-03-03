package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderer(t *testing.T) {
	// new text renderer
	tr := NewTextRenderer("text")

	// render text
	text, err := tr.Render()
	assert.NoError(t, err)
	assert.Equal(t, "text", text)

	// new json renderer
	jr := NewJSONRenderer(map[string]string{"key": "value"})

	// render json
	json, err := jr.Render()
	assert.NoError(t, err)
	assert.Equal(t, "{\"key\":\"value\"}", json)
}
