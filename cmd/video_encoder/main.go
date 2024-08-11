package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/RafaLopesMelo/go-video-encoder/internal/app/middleware"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/config/env"
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
)

func main() {
	env.Load(".env")

	mux := http.NewServeMux()
	r := router.New(mux)

	setupMiddlewares(r)
	r.Setup()

	srv := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Println("Server is running on port 3000")
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

func setupMiddlewares(r *router.Router) {
	r.Use(middleware.JSON)
	r.Use(middleware.Error)
	r.Use(middleware.Recover)
}
