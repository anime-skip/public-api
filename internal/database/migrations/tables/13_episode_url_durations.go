package tables

// AddDurationToEpisodeUrls inserts new duration column
var AddDurationToEpisodeUrls = migrateTableChange(
	"MODIFY_EPISODE_URLS_TABLE__add_duration",
	`
	ALTER TABLE public.episode_urls
		ADD duration decimal;
	`,
	`
	ALTER TABLE public.episode_urls
		DROP COLUMN duration;
	`,
)
