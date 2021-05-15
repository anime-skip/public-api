package tables

// AddDurationToEpisodes inserts new base_duration column, leaving it null (player will add missing durations)
var AddBaseDurationToEpisodes = migrateTableChange(
	"MODIFY_EPISODES_TABLE__add_base_duration",
	`
	ALTER TABLE public.episodes
		ADD base_duration decimal;
	`,
	`
	ALTER TABLE public.episodes
		DROP COLUMN base_duration;
	`,
)
