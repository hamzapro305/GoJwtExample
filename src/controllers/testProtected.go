package controllers

import "github.com/gofiber/fiber/v2"

func TestProtected(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This is a protected route",
	})
}
