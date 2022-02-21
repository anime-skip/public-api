package postgres

import "anime-skip.com/timestamps-service/internal"

type templateTimestampService struct {
	db internal.Database
}

func NewTemplateTimestampService(db internal.Database) internal.TemplateTimestampService {
	return &templateTimestampService{db}
}
