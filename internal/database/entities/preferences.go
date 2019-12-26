package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

// Preferences represents the user's settings and configuration
type Preferences struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time

	UserID           uuid.UUID `gorm:"not null;type:uuid"`
	EnableAutoSkip   bool
	EnableAutoPlay   bool
	SkipBranding     bool
	SkipIntros       bool
	SkipNewIntros    bool
	SkipRecaps       bool
	SkipFiller       bool
	SkipCanon        bool
	SkipTransitions  bool
	SkipCredits      bool
	SkipMixedCredits bool
	SkipPreview      bool
	SkipTitleCard    bool
}
