package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

// BaseEntity defines the common columns that all db structs should hold
type BaseEntity struct {
	ID              uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID uuid.UUID  `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time  `gorm:"not null"`
	UpdatedByUserID uuid.UUID  `gorm:"not null;type:uuid"`
	DeletedAt       *time.Time `gorm:""`
	DeletedByUserID *uuid.UUID `gorm:"type:uuid"`
}
