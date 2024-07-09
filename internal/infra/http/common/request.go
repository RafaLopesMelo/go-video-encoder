package common

import (
	"net/http"
	"strings"
)

type Request struct {
	Method string
	Path   string
	Params map[string]string
	Query  map[string]string
}

func (r *Request) SetParam(key string, value string) {
	r.Params[key] = value
}

func (r *Request) SetQuery(key string, value string) {
	r.Query[key] = value
}

func NewRequestFromVanilla(req *http.Request) *Request {
	return &Request{
		Method: strings.ToUpper(req.Method),
		Path:   req.URL.Path,
		Params: make(map[string]string),
		Query:  make(map[string]string),
	}
}
