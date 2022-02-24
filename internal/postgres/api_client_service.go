package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type apiClientService struct {
	db internal.Database
}

func NewAPIClientService(db internal.Database) internal.APIClientService {
	return &apiClientService{db}
}

func (s *apiClientService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]internal.APIClient, error) {
	return getAPIClientsByUserID(ctx, s.db, userID)
}

func (s *apiClientService) Create(ctx context.Context, newAPIClient internal.APIClient) (internal.APIClient, error) {
	return insertAPIClient(ctx, s.db, newAPIClient)
}

func (s *apiClientService) Update(ctx context.Context, newAPIClient internal.APIClient) (internal.APIClient, error) {
	return updateAPIClient(ctx, s.db, newAPIClient)
}

func (s *apiClientService) Delete(ctx context.Context, apiClient internal.APIClient) (internal.APIClient, error) {
	panic("apiClientService.Delete not implemented")
}
