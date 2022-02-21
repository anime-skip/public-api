package postgres

import "anime-skip.com/timestamps-service/internal"

type templateService struct {
	db internal.Database
}

func NewTemplateService(db internal.Database) internal.TemplateService {
	return &templateService{db}
}
