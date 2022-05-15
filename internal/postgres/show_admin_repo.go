package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"github.com/gofrs/uuid"
)

func deleteCascadeShowAdmin(ctx context.Context, tx internal.Tx, admin internal.ShowAdmin, deletedBy uuid.UUID) (internal.ShowAdmin, error) {
	log.V("Deleting show admin (nothing to cascade): %v", admin.ID)
	return deleteShowAdmin(ctx, tx, admin, deletedBy)
}
