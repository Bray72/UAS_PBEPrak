package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response adalah struktur standar untuk semua API response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse mengirim response sukses
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse mengirim response error
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}
