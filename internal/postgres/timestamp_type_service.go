package postgres

import "anime-skip.com/timestamps-service/internal"

type timestampTypeService struct {
	db internal.Database
}

func NewTimestampTypeService(db internal.Database) internal.TimestampTypeService {
	return &timestampTypeService{db}
}
