package ksatriya

import "net/http"

type Response struct {
	StatusCode int
	Header     http.Header
	Body       string
}

func NewResponse() *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{},
	}
}

func (res *Response) SetContentType(contentType string) {
	res.Header.Set("Content-Type", contentType)
}

func (res *Response) Write(w http.ResponseWriter) {
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)
	w.Write([]byte(res.Body))
}
