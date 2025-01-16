package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamzapro305/GoJwt/src/config"
	"github.com/hamzapro305/GoJwt/src/routes"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		return
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
