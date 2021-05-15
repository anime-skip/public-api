package entities

import "github.com/gofrs/uuid"

type TemplateTimestamp struct {
	TemplateID  uuid.UUID `gorm:"not null;type:uuid"`
	TimestampID uuid.UUID `gorm:"not null;type:uuid"`
}
