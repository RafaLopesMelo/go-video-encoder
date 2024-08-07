package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func main() {
	router.Setup()
	srv := http.Server{
		Addr: ":3000",
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()

	err := srv.Shutdown(context.TODO())
	if err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
}
