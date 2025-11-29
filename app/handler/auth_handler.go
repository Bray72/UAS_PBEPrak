package handler

import (
	"clean-arch/app/model"
	"clean-arch/app/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login handles user login
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body model.LoginRequest true "Login credentials"
// @Success      200 {object} model.LoginResponse
// @Failure      400 {object} model.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	resp, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Status:  "error",
			Message: resp.Message,
			Code:    401,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// Register handles user registration
// @Summary      User registration
// @Description  Create new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body model.RegisterRequest true "Registration data"
// @Success      201 {object} model.RegisterResponse
// @Failure      400 {object} model.ErrorResponse
// @Router       /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req := new(model.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	resp, err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: resp.Message,
			Code:    400,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetProfile handles getting current user profile
// @Summary      Get user profile
// @Description  Get authenticated user profile information
// @Tags         Auth
// @Security     Bearer
// @Produce      json
// @Success      200 {object} model.ProfileResponse
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/v1/auth/profile [get]
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	// Get user ID from token (set by middleware)
	userID := c.Locals("user_id").(string)

	resp, err := h.service.GetProfile(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Status:  "error",
			Message: resp.Message,
			Code:    404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// Logout handles user logout
// @Summary      User logout
// @Description  Logout user and invalidate token
// @Tags         Auth
// @Security     Bearer
// @Produce      json
// @Success      200 {object} model.LogoutResponse
// @Router       /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	err := h.service.Logout(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Failed to logout",
			Code:    500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.LogoutResponse{
		Status:  "success",
		Message: "Successfully logged out",
	})
}

// RefreshToken handles token refresh
// @Summary      Refresh token
// @Description  Generate new JWT token using refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body model.RefreshRequest true "Refresh token"
// @Success      200 {object} model.RefreshResponse
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	req := new(model.RefreshRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	resp, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid refresh token",
			Code:    401,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
