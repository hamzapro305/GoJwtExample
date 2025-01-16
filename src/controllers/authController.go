package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}
