package query

import (
	"encoding/json"
	"net/http"
)

type HcController struct{}

func (c *HcController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"message": "Hello World"})
}

func NewHcController() *HcController {
	return &HcController{}
}
