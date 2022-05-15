package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"github.com/gofrs/uuid"
)

type apiClientService struct {
	db internal.Database
}

func NewAPIClientService(db internal.Database) internal.APIClientService {
	return &apiClientService{db}
}

func (s *apiClientService) Get(ctx context.Context, filter internal.APIClientsFilter) (internal.APIClient, error) {
	return inTx(ctx, s.db, false, internal.ZeroAPIClient, func(tx internal.Tx) (internal.APIClient, error) {
		return findAPIClient(ctx, tx, filter)
	})
}

func (s *apiClientService) List(ctx context.Context, filter internal.APIClientsFilter) ([]internal.APIClient, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.APIClient, error) {
		return findAPIClients(ctx, tx, filter)
	})
}

func (s *apiClientService) Create(ctx context.Context, newAPIClient internal.APIClient, createdBy uuid.UUID) (internal.APIClient, error) {
	return inTx(ctx, s.db, true, internal.ZeroAPIClient, func(tx internal.Tx) (internal.APIClient, error) {
		return createAPIClient(ctx, tx, newAPIClient, createdBy)
	})
}

func (s *apiClientService) Update(ctx context.Context, newAPIClient internal.APIClient, updatedBy uuid.UUID) (internal.APIClient, error) {
	return inTx(ctx, s.db, true, internal.ZeroAPIClient, func(tx internal.Tx) (internal.APIClient, error) {
		return updateAPIClient(ctx, tx, newAPIClient, updatedBy)
	})
}

func (s *apiClientService) Delete(ctx context.Context, id string, deletedBy uuid.UUID) (internal.APIClient, error) {
	return inTx(ctx, s.db, true, internal.ZeroAPIClient, func(tx internal.Tx) (internal.APIClient, error) {
		existing, err := findAPIClient(ctx, tx, internal.APIClientsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroAPIClient, err
		}
		return deleteCascadeAPIClient(ctx, tx, existing, deletedBy)
	})
}
