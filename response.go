package ksatriya

import "net/http"

type Response struct {
	statusCode int
	header     http.Header
	body       string
}

func (res *Response) StatusCode() int {
	return res.statusCode
}
func (res *Response) SetStatusCode(val int) {
	res.statusCode = val
}

func (res *Response) Header() http.Header {
	return res.header
}
func (res *Response) SetHeader(key, val string) {
	res.header.Set(key, val)
}

func (res *Response) SetContentType(val string) {
	res.SetHeader("Content-Type", val)
}

func (res *Response) Body() string {
	return res.body
}
func (res *Response) SetBody(val string) {
	res.body = val
}

func (res *Response) WriteHeaderAndBody(w http.ResponseWriter) {
	res.WriteHeader(w)
	res.WriteBody(w)
}

func (res *Response) WriteHeader(w http.ResponseWriter) {
	for key, vals := range res.Header() {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(res.StatusCode())
}

func (res *Response) WriteBody(w http.ResponseWriter) {
	w.Write([]byte(res.Body()))
}

func NewResponse() *Response {
	return &Response{
		statusCode: http.StatusOK,
		header:     http.Header{},
	}
}
