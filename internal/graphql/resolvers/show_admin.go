package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getShowAdminsByShowId(ctx context.Context, showID *uuid.UUID) ([]*internal.ShowAdmin, error) {
	admins, err := r.ShowAdminService.List(ctx, internal.ShowAdminsFilter{
		ShowID: showID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(admins), nil
}

func (r *Resolver) getShowAdminsByUserId(ctx context.Context, userID *uuid.UUID) ([]*internal.ShowAdmin, error) {
	admins, err := r.ShowAdminService.List(ctx, internal.ShowAdminsFilter{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(admins), nil
}

// Mutations

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput internal.InputShowAdmin) (*internal.ShowAdmin, error) {
	panic(errors.NewPanicedError("TODO - show admins are disabled"))
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, id *uuid.UUID) (*internal.ShowAdmin, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.ShowAdminService.Delete(ctx, *id, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

func (r *queryResolver) FindShowAdmin(ctx context.Context, id *uuid.UUID) (*internal.ShowAdmin, error) {
	admin, err := r.ShowAdminService.Get(ctx, internal.ShowAdminsFilter{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.ShowAdmin, error) {
	return r.getShowAdminsByShowId(ctx, showID)
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID *uuid.UUID) ([]*internal.ShowAdmin, error) {
	return r.getShowAdminsByUserId(ctx, userID)
}

// Fields

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *internal.ShowAdmin) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *internal.ShowAdmin) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *internal.ShowAdmin) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *internal.ShowAdmin) (*internal.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
}

func (r *showAdminResolver) User(ctx context.Context, obj *internal.ShowAdmin) (*internal.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
