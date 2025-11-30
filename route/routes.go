package routes

import (
	"clean-arch/app/handler"
	"clean-arch/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes configures all authentication routes
func SetupAuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	authGroup := app.Group("/api/v1/auth")

	// Public routes (no authentication required)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/refresh", authHandler.RefreshToken)

	// Protected routes (authentication required)
	protected := authGroup.Use(middleware.AuthMiddleware)
	protected.Get("/profile", authHandler.GetProfile)
	protected.Post("/logout", authHandler.Logout)
}

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
			"message": "API is running",
		})
	})

	// API Routes
	SetupAuthRoutes(app, authHandler)

	// 404 handler
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"message": "Endpoint not found",
			"code": 404,
		})
	})
}
