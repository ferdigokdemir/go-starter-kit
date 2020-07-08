package routes

import (
	"fmt"
	"github.com/gofiber/fiber"
	"go-starter-kit/controllers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("Hello, 👋!")
		c.Send(msg)
	})

	categories := v1.Group("/categories")
	categories.Get("/", controllers.GetCategories)
}
