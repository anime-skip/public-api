package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

func deleteCascadeShowAdmin(ctx context.Context, tx internal.Tx, admin internal.ShowAdmin) (internal.ShowAdmin, error) {
	log.V("Deleting show admin (nothing to cascade): %v", admin.ID)
	return deleteShowAdminInTx(ctx, tx, admin)
}
