package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type APIClient struct {
	// Custom Soft Delete
	ID              string     `gorm:"primary_key"`
	CreatedAt       time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID uuid.UUID  `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time  `gorm:"not null"`
	UpdatedByUserID uuid.UUID  `gorm:"not null;type:uuid"`
	DeletedAt       *time.Time `gorm:""`
	DeletedByUserID *uuid.UUID `gorm:"type:uuid"`

	// Fields
	UserID         uuid.UUID
	AppName        string
	Description    string
	AllowedOrigins *string
	RateLimitRPM   *uint
}
