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

	j := map[string]string{}
	j["key"] = "value"

	// new json renderer
	jr := NewJSONRenderer(j)

	// render json
	assert.Equal(t, "{\"key\":\"value\"}", jr.Render())
}
