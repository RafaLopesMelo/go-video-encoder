package main

import (
	"log"
	go_http "net/http"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http"
)

func main() {
	router := http.NewRouter()

	server := go_http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Printf("[info] start http server listening 3000")
	server.ListenAndServe()
}
