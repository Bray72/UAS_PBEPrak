package service

import (
	"clean-arch/app/model"
	"clean-arch/app/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementService interface {
	GetAchievementsByUserID(userID string) ([]*model.AchievementResponse, error)
	GetAchievementByID(id string) (*model.AchievementResponse, error)
	CreateAchievement(userID string, req *model.CreateAchievementRequest) (*model.AchievementResponse, error)
	UpdateAchievement(id string, userID string, req *model.UpdateAchievementRequest) (*model.AchievementResponse, error)
	DeleteAchievement(id string, userID string) error
	SubmitAchievement(id string, userID string, req *model.SubmitAchievementRequest) (*model.AchievementResponse, error)
}

type achievementService struct {
	repo repository.AchievementRepository
}

func NewAchievementService(repo repository.AchievementRepository) AchievementService {
	return &achievementService{repo: repo}
}

func (s *achievementService) GetAchievementsByUserID(userID string) ([]*model.AchievementResponse, error) {
	achievements, err := s.repo.GetAchievementsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []*model.AchievementResponse
	for _, achievement := range achievements {
		responses = append(responses, s.achievementToResponse(achievement))
	}

	return responses, nil
}

func (s *achievementService) GetAchievementByID(id string) (*model.AchievementResponse, error) {
	achievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return nil, err
	}

	return s.achievementToResponse(achievement), nil
}

func (s *achievementService) CreateAchievement(userID string, req *model.CreateAchievementRequest) (*model.AchievementResponse, error) {
	// Validate input
	if err := s.validateCreateRequest(req); err != nil {
		return nil, err
	}

	achievement := &model.Achievement{
		ID:          primitive.NewObjectID(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Document:    req.Document,
		Status:      "draft",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdAchievement, err := s.repo.CreateAchievement(achievement)
	if err != nil {
		return nil, errors.New("failed to create achievement")
	}

	// Create reference in PostgreSQL
	pgRef := &model.AchievementPostgres{
		ID:        uuid.New(),
		UserID:    uuid.MustParse(userID),
		MongoID:   createdAchievement.ID.Hex(),
		Title:     createdAchievement.Title,
		Status:    "draft",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_ = s.repo.CreateAchievementReference(pgRef)

	return s.achievementToResponse(createdAchievement), nil
}

func (s *achievementService) UpdateAchievement(id string, userID string, req *model.UpdateAchievementRequest) (*model.AchievementResponse, error) {
	// Get existing achievement
	existingAchievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return nil, err
	}

	// Check ownership
	if existingAchievement.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	// Only draft can be updated
	if existingAchievement.Status != "draft" {
		return nil, errors.New("only draft achievement can be updated")
	}

	// Validate input
	if err := s.validateUpdateRequest(req); err != nil {
		return nil, err
	}

	// Update fields
	existingAchievement.Title = req.Title
	existingAchievement.Description = req.Description
	existingAchievement.Document = req.Document
	existingAchievement.UpdatedAt = time.Now()

	if err := s.repo.UpdateAchievement(id, existingAchievement); err != nil {
		return nil, err
	}

	// Get updated achievement
	updatedAchievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return nil, err
	}

	return s.achievementToResponse(updatedAchievement), nil
}

func (s *achievementService) DeleteAchievement(id string, userID string) error {
	// Get achievement
	achievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return err
	}

	// Check ownership
	if achievement.UserID != userID {
		return errors.New("unauthorized")
	}

	// Only draft can be deleted
	if achievement.Status != "draft" {
		return errors.New("only draft achievement can be deleted")
	}

	return s.repo.DeleteAchievement(id)
}

func (s *achievementService) SubmitAchievement(id string, userID string, req *model.SubmitAchievementRequest) (*model.AchievementResponse, error) {
	// Get achievement
	achievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return nil, err
	}

	// Check ownership
	if achievement.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	// Submit
	if err := s.repo.SubmitAchievement(id, userID); err != nil {
		return nil, err
	}

	// Update PostgreSQL status
	_ = s.repo.UpdateAchievementStatus(id, "submitted")

	// Get updated achievement
	updatedAchievement, err := s.repo.GetAchievementByID(id)
	if err != nil {
		return nil, err
	}

	return s.achievementToResponse(updatedAchievement), nil
}

func (s *achievementService) validateCreateRequest(req *model.CreateAchievementRequest) error {
	if req.Title == "" || len(req.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}

	if req.Description == "" || len(req.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}

	if req.Document == "" {
		return errors.New("document is required")
	}

	return nil
}

func (s *achievementService) validateUpdateRequest(req *model.UpdateAchievementRequest) error {
	if req.Title == "" || len(req.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}

	if req.Description == "" || len(req.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}

	if req.Document == "" {
		return errors.New("document is required")
	}

	return nil
}

func (s *achievementService) achievementToResponse(achievement *model.Achievement) *model.AchievementResponse {
	return &model.AchievementResponse{
		ID:          achievement.ID.Hex(),
		UserID:      achievement.UserID,
		Title:       achievement.Title,
		Description: achievement.Description,
		Document:    achievement.Document,
		Status:      achievement.Status,
		SubmitDate:  achievement.SubmitDate,
		CreatedAt:   achievement.CreatedAt,
		UpdatedAt:   achievement.UpdatedAt,
	}
}
