package ksatriya

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	// new root
	root := NewRoot()
	assert.Equal(t, RootPathDefault+"/*filepath", root.Path())
	assert.Equal(t, http.Dir(RootDirDefault), root.Dir())

	// set path
	root.SetPath("/public")
	assert.Equal(t, "/public/*filepath", root.Path())

	// set dir
	root.SetDir("public")
	assert.Equal(t, http.Dir("public"), root.Dir())
}
