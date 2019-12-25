package models

import "github.com/gofrs/uuid"

// Timestamp represents a point in a episode when a new section begins
type Timestamp struct {
	BaseModel
	At        float32
	TypeID    uuid.UUID
	EpisodeID uuid.UUID
}
