package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"github.com/gofrs/uuid"
)

type preferencesService struct {
	db internal.Database
}

func NewPreferencesService(db internal.Database) internal.PreferencesService {
	return &preferencesService{db}
}

func (s *preferencesService) GetByUserID(ctx context.Context, userID uuid.UUID) (internal.Preferences, error) {
	return getPreferencesByUserID(ctx, s.db, userID)
}
