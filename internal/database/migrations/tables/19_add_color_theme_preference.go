package tables

import (
	"fmt"

	"anime-skip.com/backend/internal/utils/constants"
)

// AddColorThemePreference inserts one new preference
var AddColorThemePreference = migrateTableChange(
	"MODIFY_PREFERENCES_TABLE__add_color_theme",
	fmt.Sprintf(`
	ALTER TABLE public.preferences
		ADD color_theme integer NOT NULL DEFAULT %d;
	`, constants.THEME_ANIME_SKIP_BLUE),
	`
	ALTER TABLE public.preferences
		DROP COLUMN color_theme;
	`,
)
