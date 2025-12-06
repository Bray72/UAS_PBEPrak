package service

import (
	"clean-arch/app/model"
	"clean-arch/app/repository"
	"clean-arch/utils"
	"errors"
	"github.com/google/uuid"
	
)

type AuthService interface {
	Login(req *model.LoginRequest) (*model.LoginResponse, error)
	Register(req *model.RegisterRequest) (*model.RegisterResponse, error)
	GetProfile(userID string) (*model.ProfileResponse, error)
	RefreshToken(refreshToken string) (*model.RefreshResponse, error)
	Logout(userID string) error
	ValidateToken(token string) (*model.JWTClaims, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return &model.LoginResponse{
			Status:  "error",
			Message: "Username and password are required",
		}, errors.New("invalid credentials")
	}

	// Get user from database
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil {
		return &model.LoginResponse{
			Status:  "error",
			Message: "Invalid credentials",
		}, errors.New("user not found")
	}

	// Check if user is active
	if !user.IsActive {
		return &model.LoginResponse{
			Status:  "error",
			Message: "User account is inactive",
		}, errors.New("user inactive")
	}

	// Verify password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return &model.LoginResponse{
			Status:  "error",
			Message: "Invalid credentials",
		}, errors.New("invalid password")
	}

	// Get user permissions
	permissions, err := s.repo.GetUserPermissions(user.ID.String())
	if err != nil {
		permissions = []string{}
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user, permissions)
	if err != nil {
		return &model.LoginResponse{
			Status:  "error",
			Message: "Failed to generate token",
		}, err
	}

	// Generate refresh token (same for now, in production should be different)
	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return &model.LoginResponse{
			Status:  "error",
			Message: "Failed to generate refresh token",
		}, err
	}

	return &model.LoginResponse{
		Status: "success",
		Data: &model.LoginData{
			Token:        token,
			RefreshToken: refreshToken,
			User: &model.UserResponse{
				ID:          user.ID.String(),
				Username:    user.Username,
				Email:       user.Email,
				FullName:    user.FullName,
				Permissions: permissions,
				IsActive:    user.IsActive,
				CreatedAt:   user.CreatedAt,
				UpdatedAt:   user.UpdatedAt,
			},
		},
	}, nil
}

func (s *authService) Register(req *model.RegisterRequest) (*model.RegisterResponse, error) {
	// Validate request
	if req.Username == "" || req.Email == "" || req.Password == "" || req.FullName == "" {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "All fields are required",
		}, errors.New("missing required fields")
	}

	// Check if username exists
	existingUser, _ := s.repo.GetUserByUsername(req.Username)
	if existingUser != nil {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "Username already exists",
		}, errors.New("username exists")
	}

	// Check if email exists
	existingEmail, _ := s.repo.GetUserByEmail(req.Email)
	if existingEmail != nil {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "Email already exists",
		}, errors.New("email exists")
	}

	// Validate role exists
	_, err := s.repo.GetRoleByID(req.RoleID)
	if err != nil {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "Invalid role",
		}, err
	}

	// Hash password
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "Failed to process password",
		}, err
	}

	// Create user
	roleID, _ := uuid.Parse(req.RoleID)
	newUser := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
		FullName:     req.FullName,
		RoleID:       roleID,
		IsActive:     true,
	}

	createdUser, err := s.repo.CreateUser(newUser)
	if err != nil {
		return &model.RegisterResponse{
			Status:  "error",
			Message: "Failed to create user",
		}, err
	}

	// Get permissions
	permissions, _ := s.repo.GetUserPermissions(createdUser.ID.String())

	return &model.RegisterResponse{
		Status: "success",
		Data: &model.UserResponse{
			ID:          createdUser.ID.String(),
			Username:    createdUser.Username,
			Email:       createdUser.Email,
			FullName:    createdUser.FullName,
			Permissions: permissions,
			IsActive:    createdUser.IsActive,
			CreatedAt:   createdUser.CreatedAt,
			UpdatedAt:   createdUser.UpdatedAt,
		},
		Message: "User created successfully",
	}, nil
}

func (s *authService) GetProfile(userID string) (*model.ProfileResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return &model.ProfileResponse{
			Status:  "error",
			Message: "User not found",
		}, err
	}

	permissions, _ := s.repo.GetUserPermissions(userID)

	return &model.ProfileResponse{
		Status: "success",
		Data: &model.UserResponse{
			ID:          user.ID.String(),
			Username:    user.Username,
			Email:       user.Email,
			FullName:    user.FullName,
			Permissions: permissions,
			IsActive:    user.IsActive,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
	}, nil
}

func (s *authService) RefreshToken(refreshTokenString string) (*model.RefreshResponse, error) {
	// Validate refresh token
	claims, err := utils.ValidateToken(refreshTokenString)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Get user
	user, err := s.repo.GetUserByID(claims.UserID.String())
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Get permissions
	permissions, _ := s.repo.GetUserPermissions(user.ID.String())

	// Generate new tokens
	newToken, err := utils.GenerateToken(user, permissions)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &model.RefreshResponse{
		Token:        newToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) Logout(userID string) error {
	// In a real application, you might want to:
	// 1. Invalidate tokens in Redis/cache
	// 2. Update last logout time in database
	// For now, logout is handled client-side by removing the token
	return nil
}

func (s *authService) ValidateToken(token string) (*model.JWTClaims, error) {
	return utils.ValidateToken(token)
}
