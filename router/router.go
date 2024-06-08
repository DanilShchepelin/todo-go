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
}
