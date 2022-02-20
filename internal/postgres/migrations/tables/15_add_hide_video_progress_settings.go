package tables

// AddHideTimelinePreferences inserts two new preferences
var AddHideTimelinePreferences = sqlMigration(
	"MODIFY_PREFERENCES_TABLE__add_hide_video_progress",
	`
	ALTER TABLE preferences
		ADD minimize_toolbar_when_editing boolean,
		ADD hide_timeline_when_minimized  boolean;
	`,
	`
	ALTER TABLE preferences
		DROP COLUMN minimize_toolbar_when_editing,
		DROP COLUMN hide_timeline_when_minimized;
	`,
)
