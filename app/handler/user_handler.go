package handler

import (
	"clean-arch/app/model"
	"clean-arch/app/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAllUsers lists all users
// @Summary      List all users
// @Description  Get all users with pagination (admin only)
// @Tags         Users
// @Security     Bearer
// @Produce      json
// @Success      200 {object} model.ListUsersResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Router       /api/v1/users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch users",
			Code:    500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.ListUsersResponse{
		Status:  "success",
		Message: "Users fetched successfully",
		Data:    users,
		Code:    200,
	})
}

// GetUserByID gets user detail by ID
// @Summary      Get user detail
// @Description  Get user details by ID (admin only)
// @Tags         Users
// @Security     Bearer
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} model.DetailUserResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User ID is required",
			Code:    400,
		})
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User not found",
			Code:    404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.DetailUserResponse{
		Status:  "success",
		Message: "User fetched successfully",
		Data:    *user,
		Code:    200,
	})
}

// CreateUser creates a new user
// @Summary      Create new user
// @Description  Create a new user (admin only)
// @Tags         Users
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        body body model.CreateUserRequest true "User data"
// @Success      201 {object} model.CreateUserResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Router       /api/v1/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(model.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    400,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.CreateUserResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    *user,
		Code:    201,
	})
}

// UpdateUser updates user information
// @Summary      Update user
// @Description  Update user information (admin only)
// @Tags         Users
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body model.UpdateUserRequest true "Updated user data"
// @Success      200 {object} model.UpdateUserResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User ID is required",
			Code:    400,
		})
	}

	req := new(model.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	user, err := h.service.UpdateUser(id, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.UpdateUserResponse{
		Status:  "success",
		Message: "User updated successfully",
		Data:    *user,
		Code:    200,
	})
}

// DeleteUser deletes a user
// @Summary      Delete user
// @Description  Delete a user (admin only)
// @Tags         Users
// @Security     Bearer
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} model.DeleteUserResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User ID is required",
			Code:    400,
		})
	}

	err := h.service.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User not found",
			Code:    404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.DeleteUserResponse{
		Status:  "success",
		Message: "User deleted successfully",
		Code:    200,
	})
}

// AssignRole assigns a role to user
// @Summary      Assign role to user
// @Description  Assign a role to user (admin only)
// @Tags         Users
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body model.AssignRoleRequest true "Role data"
// @Success      200 {object} model.AssignRoleResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      403 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/v1/users/{id}/role [put]
func (h *UserHandler) AssignRole(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "User ID is required",
			Code:    400,
		})
	}

	req := new(model.AssignRoleRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
			Code:    400,
		})
	}

	user, err := h.service.AssignRole(id, req)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
				Code:    404,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.AssignRoleResponse{
		Status:  "success",
		Message: "Role assigned successfully",
		Data:    *user,
		Code:    200,
	})
}
