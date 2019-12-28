package seeders

import (
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/gofrs/uuid"
)

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
