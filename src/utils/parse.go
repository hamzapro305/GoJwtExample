package utils

import "github.com/gofiber/fiber/v2"

func ParseJSON(c *fiber.Ctx) (map[string]interface{}, error) {
	var data map[string]interface{}

	// Parse JSON body into map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}

	return data, nil
}
