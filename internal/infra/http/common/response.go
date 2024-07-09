package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status  int
	body    []byte
	headers map[string]string
}

func (r *Response) Json(data interface{}) error {
	body, err := json.Marshal(data)

	if err != nil {
		return err
	}

	r.headers["Content-Type"] = "application/json"
	r.body = body
	return nil
}

func (r *Response) Status(status int) error {
	r.status = status
	return nil
}

func (r *Response) Send(w http.ResponseWriter) {
	for key, value := range r.headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(r.status)
	w.Write(r.body)
}

func NewResponse() *Response {
	return &Response{
		status:  http.StatusOK,
		headers: make(map[string]string),
	}
}
