package model

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Request DTOs
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Response DTOs
type LoginResponse struct {
	Status string         `json:"status"`
	Data   *LoginData     `json:"data"`
	Message string        `json:"message,omitempty"`
}

type LoginData struct {
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	User         *UserResponse  `json:"user"`
}

type RefreshResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ProfileResponse struct {
	Status string         `json:"status"`
	Data   *UserResponse  `json:"data"`
	Message string        `json:"message,omitempty"`
}

type LogoutResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Status string        `json:"status"`
	Data   *UserResponse `json:"data"`
	Message string       `json:"message,omitempty"`
}

type UserResponse struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	FullName    string   `json:"full_name"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JWTClaims struct {
	UserID      uuid.UUID `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Permissions []string  `json:"permissions"`
	jwt.RegisteredClaims
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
