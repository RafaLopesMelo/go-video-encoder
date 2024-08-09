package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func JSON(handler router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) any {
		res := handler(w, r)

		_, err := json.Marshal(res)

		if err != nil {
			return nil
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return nil
	}
}
