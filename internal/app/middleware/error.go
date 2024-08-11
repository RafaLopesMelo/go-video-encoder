package middleware

import (
	"net/http"

	httperror "github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/error"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func Error(h router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) any {
		res := h(w, r)
		err, ok := res.(httperror.HttpError)

		if !ok {
			return res
		}

		w.WriteHeader(err.Status)
		return err
	}
}
