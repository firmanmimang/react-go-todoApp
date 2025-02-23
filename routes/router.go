package routes

import (
	"github.com/firmanmimang/react-go-todo/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Basic route
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
	// })

	// API group for todos
	api := app.Group("/api")
	api.Get("/todos", controllers.GetTodos)
	api.Post("/todos", controllers.CreateTodo)
	api.Patch("/todos/:id", controllers.UpdateTodo)
	api.Delete("/todos/:id", controllers.DeleteTodo)
}
