package router

import "github.com/RafaLopesMelo/go-video-encoder/internal/application/queries/hc"

func SetupRouter(router *Router) {
	router.Get("/hc", hc.NewHcController())
}
