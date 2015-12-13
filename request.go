package ksatriya

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	*http.Request
}

func (req *Request) BodyJSON(v interface{}) error {
	dec := json.NewDecoder(req.Body)
	return dec.Decode(v)
}

func NewRequest(req *http.Request) *Request {
	return &Request{req}
}
