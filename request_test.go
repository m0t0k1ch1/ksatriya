package ksatriya

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Name string `json:"name"`
}

func TestRequest_BodyJSON(t *testing.T) {
	r := strings.NewReader("{\"name\":\"m0t0k1ch1\"}")
	req, err := http.NewRequest("GET", "/", r)
	assert.NoError(t, err)

	kreq := NewRequest(req)

	var u TestUser
	kreq.BodyJSON(&u)

	assert.Equal(t, "m0t0k1ch1", u.Name)
}
