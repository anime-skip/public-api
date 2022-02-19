package tables

import (
	"fmt"

	"anime-skip.com/timestamps-service/internal"
)

// AddColorThemePreference inserts one new preference
var AddColorThemePreference = sqlMigration(
	"MODIFY_PREFERENCES_TABLE__add_color_theme",
	fmt.Sprintf(`
	ALTER TABLE preferences
		ADD color_theme integer NOT NULL DEFAULT %d;
	`, internal.THEME_ANIME_SKIP_BLUE),
	`
	ALTER TABLE preferences
		DROP COLUMN color_theme;
	`,
)
