package router

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/application/query"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/hc", query.NewHcController().Handle)
}
