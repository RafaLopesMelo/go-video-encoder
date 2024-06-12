package main

import (
    "github.com/RafaLopesMelo/go-video-encoder/internal/infra/configs/env"
	log "github.com/sirupsen/logrus"
)

func main() {
    env.Load(".env")

	log.Info("Hello World!")
}
