package main

import (
	"log"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http"
)

func main() {
	server := http.NewServer()
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

	log.Printf("[info] HTTP server listening to port 3000")
}
