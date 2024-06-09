package main

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := env.Load(".env")

	if err != nil {
		panic("Could not load environment variables")
	}

	log.Info("Hello World!")
}
