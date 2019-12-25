package models

import "github.com/gofrs/uuid"

// ShowAdmin represets the link between the user and show that that user is an admin of. Admins are in charge of approving change requests
type ShowAdmin struct {
	BaseModel
	ShowID uuid.UUID
	UserID uuid.UUID
}
