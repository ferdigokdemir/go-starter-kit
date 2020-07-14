package routes

import (
	"fmt"
	"go-starter-kit/controllers"

	"github.com/gofiber/fiber"
)

// RegisterRoutes is setup routing
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("Hello, 👋!")
		c.Send(msg)
	})

	todos := v1.Group("/todos")
	todos.Get("/", controllers.GetTodos)
}
