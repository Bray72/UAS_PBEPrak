package model

import (
	
)

// CreateUserRequest represents request to create a new user
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	FullName string `json:"full_name" validate:"required"`
	RoleID   string `json:"role_id" validate:"required,uuid"`
}

// UpdateUserRequest represents request to update a user
type UpdateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

// AssignRoleRequest represents request to assign role to user
type AssignRoleRequest struct {
	RoleID string `json:"role_id" validate:"required,uuid"`
}

// RoleData represents role information in response
type RoleData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ListUsersResponse represents list of users response
type ListUsersResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []*UserResponse `json:"data"`
	Code    int            `json:"code"`
}

// DetailUserResponse represents single user response
type DetailUserResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Code    int          `json:"code"`
}

// CreateUserResponse represents create user response
type CreateUserResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Code    int          `json:"code"`
}

// UpdateUserResponse represents update user response
type UpdateUserResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Code    int          `json:"code"`
}

// DeleteUserResponse represents delete user response
type DeleteUserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// AssignRoleResponse represents assign role response
type AssignRoleResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Code    int          `json:"code"`
}

// package model

// // CreateUserRequest represents request to create a new user
// type CreateUserRequest struct {
// 	Username string `json:"username" validate:"required,min=3,max=50"`
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=6"`
// 	FullName string `json:"full_name" validate:"required"`
// 	RoleID   string `json:"role_id" validate:"required,uuid"`
// }

// // UpdateUserRequest represents request to update a user
// type UpdateUserRequest struct {
// 	Username string `json:"username" validate:"required,min=3,max=50"`
// 	Email    string `json:"email" validate:"required,email"`
// 	FullName string `json:"full_name" validate:"required"`
// 	IsActive bool   `json:"is_active"`
// }

// // AssignRoleRequest represents request to assign role to user
// type AssignRoleRequest struct {
// 	RoleID string `json:"role_id" validate:"required,uuid"`
// }

// // RoleData represents role information in response
// type RoleData struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// // ListUsersResponse represents list of users response
// type ListUsersResponse struct {
// 	Status  string         `json:"status"`
// 	Message string         `json:"message"`
// 	Data    []UserResponse `json:"data"`
// 	Code    int            `json:"code"`
// }

// // DetailUserResponse represents single user response
// type DetailUserResponse struct {
// 	Status  string       `json:"status"`
// 	Message string       `json:"message"`
// 	Data    UserResponse `json:"data"`
// 	Code    int          `json:"code"`
// }

// // CreateUserResponse represents create user response
// type CreateUserResponse struct {
// 	Status  string       `json:"status"`
// 	Message string       `json:"message"`
// 	Data    UserResponse `json:"data"`
// 	Code    int          `json:"code"`
// }

// // UpdateUserResponse represents update user response
// type UpdateUserResponse struct {
// 	Status  string       `json:"status"`
// 	Message string       `json:"message"`
// 	Data    UserResponse `json:"data"`
// 	Code    int          `json:"code"`
// }

// // DeleteUserResponse represents delete user response
// type DeleteUserResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// }

// // AssignRoleResponse represents assign role response
// type AssignRoleResponse struct {
// 	Status  string       `json:"status"`
// 	Message string       `json:"message"`
// 	Data    UserResponse `json:"data"`
// 	Code    int          `json:"code"`
// }
