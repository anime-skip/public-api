package models

import "github.com/gofrs/uuid"

// Preferences represents the user's settings and configuration
type Preferences struct {
	BaseModel
	UserID           uuid.UUID
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
