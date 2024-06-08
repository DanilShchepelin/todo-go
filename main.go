package main

import (
	"awesomeProject/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Todo-list",
		Prefork:       true,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
	})

	router.SetupRouters(app)

	log.Fatal(app.Listen(":3000"))
}
