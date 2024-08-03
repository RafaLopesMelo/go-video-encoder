package hc

import (
	"encoding/json"
	"net/http"
)

type CheckController struct{}

func (c *CheckController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"message": "Hello World"})
}

func NewCheckController() *CheckController {
	return &CheckController{}
}
