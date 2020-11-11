package tables

// AddDurationToEpisodes inserts new base_duration column with a default value
var AddBaseDurationToEpisodes = migrateTable(
	"MODIFY_EPISODES_TABLE__add_base_duration",
	"episodes",
	[]string{
		"ALTER TABLE public.episodes",
		"ADD base_duration decimal;",
	},
)
