package postgres

import "anime-skip.com/timestamps-service/internal"

type showAdminService struct {
	db internal.Database
}

func NewShowAdminService(db internal.Database) internal.ShowAdminService {
	return &showAdminService{db}
}
