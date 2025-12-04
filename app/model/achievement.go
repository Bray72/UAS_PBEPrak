package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Achievement model untuk MongoDB
type Achievement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"user_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Document    string             `bson:"document"` // URL atau file path
	Status      string             `bson:"status"`   // draft, submitted, verified, rejected
	SubmitDate  *time.Time         `bson:"submit_date,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

// AchievementPostgres model untuk referensi di PostgreSQL
type AchievementPostgres struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid"`
	MongoID       string    `gorm:"column:mongo_id"` // Reference ke MongoDB
	Title         string
	Status        string
	SubmitDate    *time.Time
	VerificationBy *uuid.UUID `gorm:"type:uuid"` // Admin yang verify
	Notes         string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (AchievementPostgres) TableName() string {
	return "achievements"
}

// Request/Response models
type CreateAchievementRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=10"`
	Document    string `json:"document" validate:"required"` // URL atau file path
}

type UpdateAchievementRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=10"`
	Document    string `json:"document" validate:"required"`
}

type SubmitAchievementRequest struct {
	Notes string `json:"notes"`
}

type AchievementResponse struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Document    string     `json:"document"`
	Status      string     `json:"status"`
	SubmitDate  *time.Time `json:"submit_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ListAchievementResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    []*AchievementResponse  `json:"data"`
}

type DetailAchievementResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    *AchievementResponse   `json:"data"`
}

type SubmitAchievementResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    *AchievementResponse   `json:"data"`
}
