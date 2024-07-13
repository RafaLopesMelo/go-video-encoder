package main

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/application/query"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Get("/hc", query.NewHcController().Handle)
	app.Listen(":3000")
}
