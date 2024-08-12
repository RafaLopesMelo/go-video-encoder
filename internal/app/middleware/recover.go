package middleware

import (
	"fmt"
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func Recover(h router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) any {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		return h(w, r)
	}
}
