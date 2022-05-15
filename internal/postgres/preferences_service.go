package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
)

type preferencesService struct {
	db internal.Database
}

func NewPreferencesService(db internal.Database) internal.PreferencesService {
	return &preferencesService{db}
}

func (s *preferencesService) Get(ctx context.Context, filter internal.PreferencesFilter) (internal.Preferences, error) {
	return inTx(ctx, s.db, false, internal.ZeroPreferences, func(tx internal.Tx) (internal.Preferences, error) {
		return findPreferences(ctx, tx, filter)
	})
}

func (s *preferencesService) Update(ctx context.Context, newPreferences internal.Preferences) (internal.Preferences, error) {
	return inTx(ctx, s.db, true, internal.ZeroPreferences, func(tx internal.Tx) (internal.Preferences, error) {
		return updatePreferences(ctx, tx, newPreferences)
	})
}
