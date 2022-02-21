package postgres

import "anime-skip.com/timestamps-service/internal"

type showService struct {
	db internal.Database
}

func NewShowService(db internal.Database) internal.ShowService {
	return &showService{db}
}
