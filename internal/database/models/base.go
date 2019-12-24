package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// BaseModel defines the common columns that all db structs should hold
type BaseModel struct {
	ID              uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID time.Time  `gorm:"not null;type:uuid"`
	UpdatedAt       *time.Time `gorm:"index"`
	UpdatedByUserID time.Time  `gorm:"not null;type:uuid"`
}

// BaseModelSoftDelete prevents the rows from being deleted, it just marks it as deleted
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `sql:"index"`
}
