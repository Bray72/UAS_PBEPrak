package model

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string     `gorm:"uniqueIndex"`
	Email        string     `gorm:"uniqueIndex"`
	PasswordHash string
	FullName     string
	IsActive     bool      `gorm:"default:true"`
	RoleID       uuid.UUID
	Role         *Role     `gorm:"foreignKey:RoleID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type Role struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"unique"`
	Description string
	Permissions []Permission   `gorm:"many2many:role_permissions"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
}

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"unique"`
	Resource    string
	Action      string
	Description string
}

func (User) TableName() string {
	return "users"
}

func (Role) TableName() string {
	return "roles"
}

func (Permission) TableName() string {
	return "permissions"
}
