package entities

import "github.com/gofrs/uuid"

type Template struct {
	BaseEntity
	ShowID          uuid.UUID `gorm:"not null;type:uuid"`
	Type            int
	Seasons         []string
	SourceEpisodeID uuid.UUID `gorm:"not null;type:uuid"`
}
