package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamzapro305/GoJwt/src/services"
)

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func Login(c *fiber.Ctx) error {
	token, err := services.GetToken(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendString(token)
}
