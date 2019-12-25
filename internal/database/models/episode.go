package models

import "github.com/gofrs/uuid"

// Episode represents data about a give episode
type Episode struct {
	BaseModel
	Season         *int
	Number         *int
	AbsoluteNumber *int
	Name           *string
	ShowID         uuid.UUID
}
