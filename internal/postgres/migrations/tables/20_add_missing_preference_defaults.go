package tables

import (
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal/postgres/migrations/sqlx_migration"
)

// AddPreferenceDefaults inserts one new preference, after switching away from GORM, they were being
// set to null
func AddMissingPreferenceDefaults() *sqlx_migration.Migration {
	type prefDefault struct {
		column       string
		defaultValue bool
	}

	prefDefaults := []prefDefault{
		{"enable_auto_skip", true},
		{"enable_auto_play", true},
		{"skip_branding", true},
		{"skip_intros", true},
		{"skip_new_intros", false},
		{"skip_mixed_intros", false},
		{"skip_recaps", true},
		{"skip_filler", true},
		{"skip_canon", false},
		{"skip_transitions", true},
		{"skip_credits", true},
		{"skip_new_credits", false},
		{"skip_mixed_credits", true},
		{"skip_preview", true},
		{"skip_title_card", true},
		{"minimize_toolbar_when_editing", false},
		{"hide_timeline_when_minimized", false},
	}

	statements := []string{}
	for _, pref := range prefDefaults {
		statement := fmt.Sprintf(
			`
				UPDATE preferences SET %s = %t WHERE %s IS NULL;
				ALTER TABLE preferences ALTER COLUMN %s SET NOT NULL;
				ALTER TABLE preferences ALTER COLUMN %s SET DEFAULT %t;
			`,
			pref.column, pref.defaultValue, pref.column,
			pref.column,
			pref.column, pref.defaultValue,
		)
		statements = append(statements, statement)
	}

	return sqlMigration(
		"MODIFY_PREFERENCES_TABLE__add_missing_defaults",
		strings.Join(statements, "\n"),
		`
		# Don't revert this change
		`,
	)
}
