package controller

import "github.com/gofiber/fiber/v2"

func MainInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"IP": c.IP()})
}
