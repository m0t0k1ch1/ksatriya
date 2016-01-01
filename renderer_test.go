package ksatriya

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderer(t *testing.T) {
	// new text renderer
	tr := NewTextRenderer("text")

	// render text
	assert.Equal(t, "text", tr.Render())

	// new json renderer
	jr := NewJSONRenderer(map[string]string{"key": "value"})

	// render json
	assert.Equal(t, "{\"key\":\"value\"}", jr.Render())
}
