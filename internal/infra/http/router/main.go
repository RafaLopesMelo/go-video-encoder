package router

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/application/query"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(app *fiber.App) {
	app.Get("/hc", query.NewHcController().Handle)
}
