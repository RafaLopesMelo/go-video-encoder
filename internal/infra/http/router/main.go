package router

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/usecase/hc"
)

func Setup() {
	http.HandleFunc("GET /hc", hc.NewCheckController().Handle)
}
