package entities

import "github.com/gofrs/uuid"

// Timestamp represents a point in a episode when a new section begins
type Timestamp struct {
	BaseEntity
	At        float64
	Source    int       `gorm:"not null"`
	TypeID    uuid.UUID `gorm:"not null;type:uuid"`
	EpisodeID uuid.UUID `gorm:"not null;type:uuid"`
}
