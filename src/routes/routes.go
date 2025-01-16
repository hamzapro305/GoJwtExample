package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamzapro305/GoJwt/src/controllers"
	"github.com/hamzapro305/GoJwt/src/services"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Auth Routes
	auth := v1.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	// Protection Middleware
	protected := v1.Group("/protected")
	protected.Use(services.ProtectedRoute())

	// Protected Routes
	protected.Get("/test", controllers.TestProtected)
}
