package routes

import (
	"go-nonce/middleware"
	"go-nonce/src/controllers"

	"github.com/gofiber/fiber/v2"
)

type fiberRoutes struct {
	exampleController controllers.ExampleController
}

func (r fiberRoutes) Install(app *fiber.App) {
	route := app.Group("/api", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Get("/example", middleware.ValidateNonce, r.exampleController.ExampleController)

	// Nonce generation endpoint
	route.Get("/nonce", middleware.GenerateNonce)
}

func NewFiberRoutes(
	exampleController controllers.ExampleController,
) fiberRoutes {
	return fiberRoutes{
		exampleController: exampleController,
	}
}
