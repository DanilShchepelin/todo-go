package router

import (
	"awesomeProject/controller"
	_ "github.com/DanilShchepelin/todo-go/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// SetupRouters @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func SetupRouters(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault) // default

	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
	api := app.Group("/api", logger.New())
	api.Get("/", controller.Hello)

	system := api.Group("/system", logger.New())
	system.Get("/", controller.MainInfo)

	todoItems := api.Group("/todo", logger.New())
	todoItems.Get("/", controller.GetAllTodoItems)
	todoItems.Get("/:id", controller.GetTodoItemById)
	todoItems.Post("/", controller.CreateTodoItem)
	todoItems.Delete("/:id", controller.DeleteTodoItem)
}
