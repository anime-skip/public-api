package postgres

import "anime-skip.com/timestamps-service/internal"

type timestampService struct {
	db internal.Database
}

func NewTimestampService(db internal.Database) internal.TimestampTypeService {
	return &timestampTypeService{db}
}
