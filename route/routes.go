package routes

import (
	"clean-arch/app/handler"
	"clean-arch/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes configures all authentication routes
func SetupAuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	authGroup := app.Group("/api/auth")

	// Public routes (no authentication required)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/refresh", authHandler.RefreshToken)

	// Protected routes (authentication required)
	protected := authGroup.Use(middleware.AuthMiddleware)
	protected.Get("/profile", authHandler.GetProfile)
	protected.Post("/logout", authHandler.Logout)
}

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	usersGroup := app.Group("/api/users")

	// All user routes require authentication and admin role
	usersGroup.Use(middleware.AuthMiddleware)
	usersGroup.Use(middleware.RoleMiddleware("admin"))

	// User management endpoints
	usersGroup.Get("/", userHandler.GetAllUsers)
	usersGroup.Get("/:id", userHandler.GetUserByID)
	usersGroup.Post("/", userHandler.CreateUser)
	usersGroup.Put("/:id", userHandler.UpdateUser)
	usersGroup.Delete("/:id", userHandler.DeleteUser)
	usersGroup.Put("/:id/role", userHandler.AssignRole)
}

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler, userHandler *handler.UserHandler) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
		})
	})

	// API Routes
	SetupAuthRoutes(app, authHandler)
	SetupUserRoutes(app, userHandler)

	// 404 handler
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Endpoint not found",
			"code":    404,
		})
	})
}
