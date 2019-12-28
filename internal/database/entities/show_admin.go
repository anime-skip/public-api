package entities

import "github.com/gofrs/uuid"

// ShowAdmin represets the link between the user and show that that user is an admin of. Admins are in charge of approving change requests
type ShowAdmin struct {
	BaseEntity
	ShowID uuid.UUID `gorm:"not null;type:uuid"`
	UserID uuid.UUID `gorm:"not null;type:uuid"`
}
