package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/utils"
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

func (p *preferencesService) NewDefault(ctx context.Context, userID uuid.UUID) internal.Preferences {
	return internal.Preferences{
		ID:     utils.RandomID(),
		UserID: userID,

		EnableAutoSkip:             true,
		EnableAutoPlay:             true,
		MinimizeToolbarWhenEditing: false,
		HideTimelineWhenMinimized:  false,
		ColorTheme:                 internal.THEME_ANIME_SKIP_BLUE,

		SkipBranding:     true,
		SkipIntros:       true,
		SkipNewIntros:    false,
		SkipMixedIntros:  false,
		SkipRecaps:       true,
		SkipFiller:       true,
		SkipCanon:        false,
		SkipTransitions:  true,
		SkipCredits:      true,
		SkipNewCredits:   false,
		SkipMixedCredits: false,
		SkipPreview:      true,
		SkipTitleCard:    true,
	}
}

func (p *preferencesService) CreateInTx(ctx context.Context, tx internal.Tx, newPreferences internal.Preferences) (internal.Preferences, error) {
	return insertPreferencesInTx(ctx, tx, newPreferences)
}

func (p *preferencesService) Update(ctx context.Context, newPreferences internal.Preferences) (internal.Preferences, error) {
	return updatePreferences(ctx, p.db, newPreferences)
}
