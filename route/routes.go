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
