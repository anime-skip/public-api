package tables

// AddHideTimelinePreferences inserts two new preferences
var AddHideTimelinePreferences = migrateTableChange(
	"MODIFY_PREFERENCES_TABLE__add_hide_video_progress",
	`
	ALTER TABLE public.preferences
		ADD minimize_toolbar_when_editing boolean,
		ADD hide_timeline_when_minimized boolean;
	`,
	`
	ALTER TABLE public.preferences
		DROP COLUMN minimize_toolbar_when_editing,
		DROP COLUMN hide_timeline_when_minimized;
	`,
)
