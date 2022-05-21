package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"github.com/gofrs/uuid"
)

func findShowAdmins(ctx context.Context, tx internal.Tx, filter internal.ShowAdminsFilter) ([]internal.ShowAdmin, error) {
	return []internal.ShowAdmin{}, internal.NewNotImplemented("findShowAdmins")
}

func findShowAdmin(ctx context.Context, tx internal.Tx, filter internal.ShowAdminsFilter) (internal.ShowAdmin, error) {
	all, err := findShowAdmins(ctx, tx, filter)
	if err != nil {
		return internal.ZeroShowAdmin, err
	} else if len(all) == 0 {
		return internal.ZeroShowAdmin, internal.NewNotFound("Show admin", "findShowAdmin")
	}
	return all[0], nil
}

func createShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, createdBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("createShowAdmin")
}

func updateShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, updatedBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("updateShowAdmin")
}

func deleteShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, deletedBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("deleteShowAdmin")
}

func deleteCascadeShowAdmin(ctx context.Context, tx internal.Tx, admin internal.ShowAdmin, deletedBy uuid.UUID) (internal.ShowAdmin, error) {
	log.V("Deleting show admin (nothing to cascade): %v", admin.ID)
	return deleteShowAdmin(ctx, tx, admin, deletedBy)
}
