package postgres

import "anime-skip.com/timestamps-service/internal"

type apiClientService struct {
	db internal.Database
}

func NewAPIClientService(db internal.Database) internal.APIClientService {
	return &apiClientService{db}
}
