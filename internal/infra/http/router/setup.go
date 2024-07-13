package router

import "github.com/RafaLopesMelo/go-video-encoder/internal/application/query"

func SetupRouter(router *Router) {
	router.Get("/hc", query.NewHcController())
}
