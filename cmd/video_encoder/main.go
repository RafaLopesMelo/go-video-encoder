package main

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
	"net/http"
)

func main() {
	router.Setup()
	http.ListenAndServe(":3000", nil)
}
