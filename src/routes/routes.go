package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamzapro305/GoJwt/src/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Get("/register", controllers.Register)
	auth.Get("/login", controllers.Login)
}
