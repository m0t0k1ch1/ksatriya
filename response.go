package ksatriya

import "net/http"

type Response struct {
	StatusCode int
	Header     http.Header
	Result     Result
}

func (res *Response) Write(ctx *Context, w http.ResponseWriter) {
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)
	res.Result.Apply(ctx, w)
}
