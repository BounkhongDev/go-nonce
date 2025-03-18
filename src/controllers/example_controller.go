package controllers

import (
"go-nonce/src/services"
	"github.com/gofiber/fiber/v2")

type ExampleController interface{
	ExampleController(ctx *fiber.Ctx) error
}

type exampleController struct {
serviceExample services.ExampleService
}

func NewExampleController(
serviceExample services.ExampleService,
//services
) ExampleController {
	return &exampleController{
serviceExample :serviceExample,
//services
}
}
func (c *exampleController) ExampleController(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
			"message": "pong",
	})
}