package hc

import (
	"net/http"
)

type CheckController struct{}

func (c *CheckController) Handle(w http.ResponseWriter, r *http.Request) any {
	return map[string]any{"message": "Hello World"}
}

func NewCheckController() *CheckController {
	return &CheckController{}
}
