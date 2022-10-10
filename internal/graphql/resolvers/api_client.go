package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
)

// Helpers

// Mutations

func (r *mutationResolver) CreateAPIClient(ctx context.Context, client internal.CreateAPIClient) (*internal.APIClient, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newAPIClient := internal.APIClient{
		UserID: &auth.UserID,
	}
	mappers.ApplyCreateAPIClient(client, &newAPIClient)
	created, err := r.APIClientService.Create(ctx, newAPIClient, auth.UserID)

	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *mutationResolver) UpdateAPIClient(ctx context.Context, id string, changes map[string]any) (*internal.APIClient, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	isAdmin := auth.Role == internal.RoleAdmin || auth.Role == internal.RoleDev
	if _, ok := changes["rateLimitRpm"]; ok && !isAdmin {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: "You must be an admin to change a client's rate limit",
			Op:      "UpdateAPIClient",
		}
	}

	newAPIClient, err := r.APIClientService.Get(ctx, internal.APIClientsFilter{
		ID:     &id,
		UserID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}
	utils.ApplyChanges(changes, &newAPIClient)
	updated, err := r.APIClientService.Update(ctx, newAPIClient, auth.UserID)

	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *mutationResolver) DeleteAPIClient(ctx context.Context, id string) (*internal.APIClient, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.APIClientService.Delete(ctx, id, auth.UserID, auth.UserID)
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}

// Queries

// FindAPIClient implements graphql.QueryResolver
func (r *queryResolver) FindAPIClient(ctx context.Context, id string) (*internal.APIClient, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	client, err := r.APIClientService.Get(ctx, internal.APIClientsFilter{
		ID:     &id,
		UserID: &auth.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// MyAPIClients implements graphql.QueryResolver
func (r *queryResolver) MyAPIClients(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*internal.APIClient, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	clients, err := r.APIClientService.List(ctx, internal.APIClientsFilter{
		UserID:       &auth.UserID,
		NameContains: search,
		Sort:         utils.ValueOr(sort, "ASC"),
		Pagination: &internal.Pagination{
			Limit:  utils.ValueOr(limit, 10),
			Offset: utils.ValueOr(offset, 0),
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.PtrSlice(clients), nil
}

// Fields

func (r *apiClientResolver) User(ctx context.Context, obj *internal.APIClient) (*internal.User, error) {
	return r.getUserById(ctx, obj.UserID)
}

func (r *apiClientResolver) CreatedBy(ctx context.Context, obj *internal.APIClient) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *apiClientResolver) UpdatedBy(ctx context.Context, obj *internal.APIClient) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *apiClientResolver) DeletedBy(ctx context.Context, obj *internal.APIClient) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
