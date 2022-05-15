package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/utils"
)

// Helpers

// Mutations

func (r *mutationResolver) CreateAPIClient(ctx context.Context, client internal.CreateAPIClient) (*internal.APIClient, error) {
	panic("unimplemented")
}

func (r *mutationResolver) DeleteAPIClient(ctx context.Context, id string) (*internal.APIClient, error) {
	panic("unimplemented")
}

func (r *mutationResolver) UpdateAPIClient(ctx context.Context, id string, changes internal.APIClientChanges) (*internal.APIClient, error) {
	panic("unimplemented")
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
			Limit:  utils.ValueOr(offset, 10),
			Offset: utils.ValueOr(offset, 0),
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.PtrSlice(clients), nil
}

// Fields
