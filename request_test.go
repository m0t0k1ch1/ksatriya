package ksatriya

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyTestUser struct {
	Name string `json:"name"`
}

func TestRequest(t *testing.T) {
	r := strings.NewReader("{\"name\":\"m0t0k1ch1\"}")
	req, err := http.NewRequest("GET", "/", r)
	assert.NoError(t, err)

	// new request
	kreq := NewRequest(req)

	var u MyTestUser

	// decode body json
	kreq.BodyJSON(&u)
	assert.Equal(t, "m0t0k1ch1", u.Name)
}
