package http

import (
	"net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func NewServer() *http.Server {
	r := router.NewRouter()
	router.SetupRouter(r)

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	return server
}
