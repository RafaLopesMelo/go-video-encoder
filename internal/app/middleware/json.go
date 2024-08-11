package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func JSON(h router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) any {
		res := h(w, r)

		_, err := json.Marshal(res)

		if err != nil {
			return res
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return res
	}
}
