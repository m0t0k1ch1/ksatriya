package ksatriya

import "net/http"

type Res interface {
	StatusCode() int
	Header() http.Header
	Result() Result
	Writer() http.ResponseWriter

	SetStatusCode(statusCode int)
	SetResult(result Result)
	SetContentType(contentType string)

	Write(ctx Ctx)
}

type Response struct {
	statusCode int
	header     http.Header
	result     Result
	writer     http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		statusCode: http.StatusOK,
		header:     http.Header{},
		writer:     w,
	}
}

func (res *Response) StatusCode() int {
	return res.statusCode
}

func (res *Response) Header() http.Header {
	return res.header
}

func (res *Response) Result() Result {
	return res.result
}

func (res *Response) Writer() http.ResponseWriter {
	return res.writer
}

func (res *Response) SetStatusCode(statusCode int) {
	res.statusCode = statusCode
}

func (res *Response) SetResult(result Result) {
	res.result = result
}

func (res *Response) SetContentType(contentType string) {
	res.Header().Set("Content-Type", contentType)
}

func (res *Response) Write(ctx Ctx) {
	w := res.Writer()
	for key, values := range res.Header() {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode())
	res.Result().Apply(ctx, w)
}
