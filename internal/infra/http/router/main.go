package router

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/query"
)

func Setup() {
	http.HandleFunc("GET /hc", query.NewHcController().Handle)
}
