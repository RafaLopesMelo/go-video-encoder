package query

import "github.com/gofiber/fiber/v3"

type HcController struct{}

func (c *HcController) Handle(ctx fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "OK",
	})
}

func NewHcController() *HcController {
	return &HcController{}
}
