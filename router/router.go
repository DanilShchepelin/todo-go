package router

import (
	"awesomeProject/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", controller.Hello)

	system := api.Group("/system", logger.New())
	system.Get("/", controller.MainInfo)

	todoItems := api.Group("/todo", logger.New())
	todoItems.Get("/", controller.GetAllTodoItems)
	todoItems.Get("/:id", controller.GetTodoItemById)
	todoItems.Post("/", controller.CreateTodoItem)
	todoItems.Delete("/:id", controller.DeleteTodoItem)

	users := api.Group("/users", logger.New())
	users.Get("/", controller.GetAllUsers)
	users.Get("/:id", controller.GetUserById)
	users.Post("/", controller.CreateUser)
	users.Delete("/:id", controller.DeleteUser)
}
