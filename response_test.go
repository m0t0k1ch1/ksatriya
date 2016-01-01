package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	// new response
	res := NewResponse()
	assert.Equal(t, http.StatusOK, res.StatusCode())
	assert.Equal(t, http.Header{}, res.Header())
	assert.Empty(t, res.Body())

	// set status code
	res.SetStatusCode(http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode())

	// set header
	res.SetHeader("X-Ksatriya-User", "m0t0k1ch1")
	assert.Equal(t, "m0t0k1ch1", res.Header().Get("X-Ksatriya-User"))

	// set content type
	res.SetContentType("text/plain")
	assert.Equal(t, "text/plain", res.Header().Get("Content-Type"))

	// set body
	res.SetBody("error")
	assert.Equal(t, "error", res.Body())

	rec := httptest.NewRecorder()

	// write
	res.Write(rec)
	header := rec.Header()
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "m0t0k1ch1", header.Get("X-Ksatriya-User"))
	assert.Equal(t, "text/plain", header.Get("Content-Type"))
	assert.Equal(t, "error", rec.Body.String())
}
