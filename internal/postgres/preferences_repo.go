package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
)

func createPreferences(ctx context.Context, tx internal.Tx, preferences internal.Preferences) (internal.Preferences, error) {
	id, err := utils.RandomID()
	if err != nil {
		return preferences, err
	}
	preferences.ID = id
	preferences.CreatedAt = *now()
	preferences.UpdatedAt = *now()
	if err != nil {
		return preferences, err
	}

	sql, args := sqlbuilder.Insert("preferences", map[string]any{
		"id":                            preferences.ID,
		"created_at":                    preferences.CreatedAt,
		"updated_at":                    preferences.UpdatedAt,
		"deleted_at":                    preferences.DeletedAt,
		"user_id":                       preferences.UserID,
		"enable_auto_skip":              preferences.EnableAutoSkip,
		"enable_auto_play":              preferences.EnableAutoPlay,
		"minimize_toolbar_when_editing": preferences.MinimizeToolbarWhenEditing,
		"hide_timeline_when_minimized":  preferences.HideTimelineWhenMinimized,
		"color_theme":                   mappers.ToColorThemeInt(preferences.ColorTheme),
		"skip_branding":                 preferences.SkipBranding,
		"skip_intros":                   preferences.SkipIntros,
		"skip_new_intros":               preferences.SkipNewIntros,
		"skip_mixed_intros":             preferences.SkipMixedIntros,
		"skip_recaps":                   preferences.SkipRecaps,
		"skip_filler":                   preferences.SkipFiller,
		"skip_canon":                    preferences.SkipCanon,
		"skip_transitions":              preferences.SkipTransitions,
		"skip_credits":                  preferences.SkipCredits,
		"skip_new_credits":              preferences.SkipNewCredits,
		"skip_mixed_credits":            preferences.SkipMixedCredits,
		"skip_preview":                  preferences.SkipPreview,
		"skip_title_card":               preferences.SkipTitleCard,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return preferences, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "Preferences with the generated ID already exists, try again",
			Op:      "createAPIClient",
			Err:     err,
		}
	} else if err != nil {
		return preferences, internal.SQLFailure("findPreferencess", err)
	}
	return preferences, nil
}
