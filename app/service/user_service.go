package service

import (
	"clean-arch/app/model"
	"clean-arch/app/repository"
	"clean-arch/utils"
	"errors"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

type UserService interface {
	GetAllUsers() ([]*model.UserResponse, error)
	GetUserByID(id string) (*model.UserResponse, error)
	CreateUser(req *model.CreateUserRequest) (*model.UserResponse, error)
	UpdateUser(id string, req *model.UpdateUserRequest) (*model.UserResponse, error)
	DeleteUser(id string) error
	AssignRole(userID string, req *model.AssignRoleRequest) (*model.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]*model.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var responses []*model.UserResponse
	for _, user := range users {
		responses = append(responses, s.userToResponse(user))
	}

	return responses, nil
}

func (s *userService) GetUserByID(id string) (*model.UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

func (s *userService) CreateUser(req *model.CreateUserRequest) (*model.UserResponse, error) {
	// Validate input
	if err := s.validateCreateUserRequest(req); err != nil {
		return nil, err
	}

	// Check if username already exists
	existingUser, _ := s.repo.GetUserByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// Parse role ID
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return nil, errors.New("invalid role id format")
	}

	// Create user
	user := &model.User{
		Username:     req.Username,
		Email:        strings.ToLower(req.Email),
		PasswordHash: hashedPassword,
		FullName:     req.FullName,
		RoleID:       roleID,
		IsActive:     true,
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Get user with role data
	userWithRole, err := s.repo.GetUserByID(createdUser.ID.String())
	if err != nil {
		return nil, err
	}

	return s.userToResponse(userWithRole), nil
}

func (s *userService) UpdateUser(id string, req *model.UpdateUserRequest) (*model.UserResponse, error) {
	// Get existing user
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Validate input
	if err := s.validateUpdateUserRequest(req); err != nil {
		return nil, err
	}

	// Update fields
	existingUser.Username = req.Username
	existingUser.Email = strings.ToLower(req.Email)
	existingUser.FullName = req.FullName
	existingUser.IsActive = req.IsActive

	// Update user
	if err := s.repo.UpdateUser(existingUser); err != nil {
		return nil, errors.New("failed to update user")
	}

	// Get updated user with role data
	updatedUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return s.userToResponse(updatedUser), nil
}

func (s *userService) DeleteUser(id string) error {
	// Check if user exists
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(id)
}

func (s *userService) AssignRole(userID string, req *model.AssignRoleRequest) (*model.UserResponse, error) {
	// Check if user exists
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Validate role ID format
	_, err = uuid.Parse(req.RoleID)
	if err != nil {
		return nil, errors.New("invalid role id format")
	}

	// Update user role
	if err := s.repo.UpdateUserRole(userID, req.RoleID); err != nil {
		return nil, err
	}

	// Get updated user
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

func (s *userService) validateCreateUserRequest(req *model.CreateUserRequest) error {
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

func (s *userService) validateUpdateUserRequest(req *model.UpdateUserRequest) error {
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

func (s *userService) userToResponse(user *model.User) *model.UserResponse {
	resp := &model.UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return resp
}