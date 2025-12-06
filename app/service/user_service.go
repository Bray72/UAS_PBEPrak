package service

import (
	"clean-arch/app/model"
	"clean-arch/app/repository"
	"clean-arch/utils"
	"database/sql"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersService(c *fiber.Ctx, db *sql.DB) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "created_at")
	order := c.Query("order", "desc")
	search := c.Query("search", "")

	// Validasi dan sanitasi parameters
	sortByWhitelist := map[string]bool{
		"id": true, "username": true, "email": true, "full_name": true, "created_at": true,
	}
	if !sortByWhitelist[sortBy] {
		sortBy = "created_at"
	}
	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	offset := (page - 1) * limit

	users, err := repository.GetAllUsersWithPagination(db, search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch users: " + err.Error(),
			"code":    500,
		})
	}

	total, err := repository.CountUsers(db, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to count users: " + err.Error(),
			"code":    500,
		})
	}

	totalPages := (total + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Users fetched successfully",
		"data":    users,
		"meta": fiber.Map{
			"page":   page,
			"limit":  limit,
			"total":  total,
			"pages":  totalPages,
			"sortBy": sortBy,
			"order":  order,
			"search": search,
		},
		"code": 200,
	})
}

func GetUserByIDService(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID is required",
			"code":    400,
		})
	}

	user, err := repository.GetUserByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"code":    404,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch user: " + err.Error(),
			"code":    500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User fetched successfully",
		"data":    user,
		"code":    200,
	})
}

func CreateUserService(c *fiber.Ctx, db *sql.DB) error {
	req := new(model.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"code":    400,
		})
	}

	// Validasi input
	if err := validateCreateUserRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"code":    400,
		})
	}

	// Cek username sudah ada
	existing, _ := repository.GetUserByUsername(db, req.Username)
	if existing != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Username already exists",
			"code":    400,
		})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to process password",
			"code":    500,
		})
	}

	user, err := repository.CreateUser(db, req, hashedPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user: " + err.Error(),
			"code":    500,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"data":    user,
		"code":    201,
	})
}

func UpdateUserService(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID is required",
			"code":    400,
		})
	}

	req := new(model.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"code":    400,
		})
	}

	// Validasi input
	if err := validateUpdateUserRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"code":    400,
		})
	}

	user, err := repository.UpdateUser(db, id, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"code":    404,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update user: " + err.Error(),
			"code":    500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User updated successfully",
		"data":    user,
		"code":    200,
	})
}

func DeleteUserService(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID is required",
			"code":    400,
		})
	}

	err := repository.DeleteUser(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"code":    404,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete user: " + err.Error(),
			"code":    500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
		"code":    200,
	})
}

func AssignRoleService(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "User ID is required",
			"code":    400,
		})
	}

	req := new(model.AssignRoleRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"code":    400,
		})
	}

	if req.RoleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Role ID is required",
			"code":    400,
		})
	}

	user, err := repository.AssignRole(db, id, req.RoleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
				"code":    404,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Role assigned successfully",
		"data":    user,
		"code":    200,
	})
}

func validateCreateUserRequest(req *model.CreateUserRequest) error {
	if req.Username == "" || len(req.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if req.Email == "" || !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}
	if req.Password == "" || len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if req.FullName == "" {
		return errors.New("full name is required")
	}
	if req.RoleID == "" {
		return errors.New("role id is required")
	}
	return nil
}

func validateUpdateUserRequest(req *model.UpdateUserRequest) error {
	if req.Username == "" || len(req.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if req.Email == "" || !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}
	if req.FullName == "" {
		return errors.New("full name is required")
	}
	return nil
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
