package entities

import "github.com/gofrs/uuid"

// Episode represents data about a give episode
type Episode struct {
	BaseEntity
	Season         *int
	Number         *int
	AbsoluteNumber *int
	Name           *string
	ShowID         uuid.UUID `gorm:"not null;type:uuid"`
}
