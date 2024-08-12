package middleware

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func CORS(h router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) any {
		res := h(w, r)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return res
		}

		return res
	}
}
