package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
)

func findPreferences(ctx context.Context, tx internal.Tx, filter internal.PreferencesFilter) (internal.Preferences, error) {
	var scanned internal.Preferences
	var scannedColorTheme int
	query := sqlbuilder.Select("preferences", map[string]any{
		"id":                            &scanned.ID,
		"created_at":                    &scanned.CreatedAt,
		"updated_at":                    &scanned.UpdatedAt,
		"deleted_at":                    &scanned.DeletedAt,
		"user_id":                       &scanned.UserID,
		"enable_auto_skip":              &scanned.EnableAutoSkip,
		"enable_auto_play":              &scanned.EnableAutoPlay,
		"minimize_toolbar_when_editing": &scanned.MinimizeToolbarWhenEditing,
		"hide_timeline_when_minimized":  &scanned.HideTimelineWhenMinimized,
		"color_theme":                   &scannedColorTheme,
		"skip_branding":                 &scanned.SkipBranding,
		"skip_intros":                   &scanned.SkipIntros,
		"skip_new_intros":               &scanned.SkipNewIntros,
		"skip_mixed_intros":             &scanned.SkipMixedIntros,
		"skip_recaps":                   &scanned.SkipRecaps,
		"skip_filler":                   &scanned.SkipFiller,
		"skip_canon":                    &scanned.SkipCanon,
		"skip_transitions":              &scanned.SkipTransitions,
		"skip_credits":                  &scanned.SkipCredits,
		"skip_new_credits":              &scanned.SkipNewCredits,
		"skip_mixed_credits":            &scanned.SkipMixedCredits,
		"skip_preview":                  &scanned.SkipPreview,
		"skip_title_card":               &scanned.SkipTitleCard,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}

	sql, args := query.ToSQL()
	row := tx.QueryRowContext(ctx, sql, args...)

	dest := query.ScanDest()
	err := row.Scan(dest...)
	if err != nil {
		return internal.ZeroPreferences, internal.SQLFailure("findPreferences", err)
	}
	return scanned, nil
}

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
			Message: "Preferences with the generated UUID already exists, try again",
			Op:      "createAPIClient",
			Err:     err,
		}
	} else if err != nil {
		return preferences, internal.SQLFailure("findPreferences", err)
	}
	return preferences, nil
}

func updatePreferences(ctx context.Context, tx internal.Tx, preferences internal.Preferences) (internal.Preferences, error) {
	preferences.UpdatedAt = *now()

	sql, args := sqlbuilder.Update("preferences", preferences.ID, map[string]any{
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

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return preferences, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Preferences",
			Op:      "updatePreferences",
			Err:     err,
		}
	}

	return preferences, nil
}
