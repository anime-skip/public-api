package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

// EpisodeURL represents one of the the URLs that this episode can be found at
type EpisodeURL struct {
	URL             string    `gorm:"primary_key"`
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID uuid.UUID `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time `gorm:"not null"`
	UpdatedByUserID uuid.UUID `gorm:"not null;type:uuid"`

	Source           int       `gorm:"not null"`
	Duration         *float64  `gorm:"type:decimal"`
	TimestampsOffset *float64  `gorm:"type:decimal"`
	EpisodeID        uuid.UUID `gorm:"not null;type:uuid"`
}
