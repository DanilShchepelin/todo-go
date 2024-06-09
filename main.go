package main

import (
	"awesomeProject/database"
	_ "awesomeProject/docs"
	"awesomeProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

//	@title			Список заданий
//	@version		1.0
//	@description	Тестовое приложение для списка заданий
func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Todo-list",
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	database.ConnectDB()

	router.SetupRouters(app)

	log.Fatal(app.Listen(":3000"))
}
