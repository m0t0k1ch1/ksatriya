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

func (res *Response) SetContentType(val string) {
	res.Header.Set("Content-Type", val)
}

func (res *Response) Write(w http.ResponseWriter) {
	for key, vals := range res.Header {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(res.StatusCode)
	w.Write([]byte(res.Body))
}
