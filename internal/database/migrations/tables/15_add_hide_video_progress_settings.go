package tables

// AddHideTimelinePreferences inserts new timestamps_offset column
var AddHideTimelinePreferences = migrateTableChange(
	"MODIFY_PREFERENCES_TABLE__add_hide_video_progress",
	[]string{
		"ALTER TABLE public.preferences",
		"ADD minimize_toolbar_when_editing boolean,",
		"ADD hide_timeline_when_minimized boolean;",
	},
	[]string{
		"ALTER TABLE public.preferences",
		"DROP COLUMN minimize_toolbar_when_editing,",
		"DROP COLUMN hide_timeline_when_minimized;",
	},
)
