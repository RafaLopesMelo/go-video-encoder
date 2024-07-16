package main

import (
	"github.com/RafaLopesMelo/go-video-encoder/internal/infra/http/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	router.Setup(app)
	app.Listen(":3000")
}
