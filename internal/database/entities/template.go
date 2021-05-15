package entities

import (
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

type Template struct {
	BaseEntity
	ShowID          uuid.UUID `gorm:"not null;type:uuid"`
	Type            int
	Seasons         pq.StringArray `gorm:"type:text[]"`
	SourceEpisodeID uuid.UUID      `gorm:"not null;type:uuid"`
}
