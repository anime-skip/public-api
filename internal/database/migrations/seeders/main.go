package seeders

import (
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/gofrs/uuid"
)

var adminUUID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000000")
var now = time.Now()

func basicEntity(id uuid.UUID) entities.BaseEntity {
	return entities.BaseEntity{
		ID:              id,
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		DeletedAt:       nil,
		DeletedByUserID: nil,
	}
}
