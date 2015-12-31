package ksatriya

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	res := NewResponse()

	assert.Equal(t, http.StatusOK, res.StatusCode())
	assert.Equal(t, http.Header{}, res.Header())
	assert.Zero(t, res.Body())

	res.SetStatusCode(http.StatusInternalServerError)
	res.SetHeader("X-Ksatriya-User", "m0t0k1ch1")
	res.SetContentType("text/plain")
	res.SetBody("error")

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode())
	assert.Equal(t, "m0t0k1ch1", res.Header().Get("X-Ksatriya-User"))
	assert.Equal(t, "text/plain", res.Header().Get("Content-Type"))
	assert.Equal(t, "error", res.Body())

	rec := httptest.NewRecorder()
	res.Write(rec)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "m0t0k1ch1", rec.Header().Get("X-Ksatriya-User"))
	assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
	assert.Equal(t, "error", rec.Body.String())
}
