package tables

// AddDurationToEpisodes inserts new base_duration column with a default value
var AddBaseDurationToEpisodes = migrateTableChange(
	"MODIFY_EPISODES_TABLE__add_base_duration",
	[]string{
		"ALTER TABLE public.episodes",
		"ADD base_duration decimal;",
	},
	[]string{
		"ALTER TABLE public.episodes",
		"DROP COLUMN base_duration;",
	},
)
