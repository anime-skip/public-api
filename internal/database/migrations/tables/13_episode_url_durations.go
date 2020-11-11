package tables

// AddDurationToEpisodeUrls inserts new duration column
var AddDurationToEpisodeUrls = migrateTable(
	"MODIFY_EPISODE_URLS_TABLE__add_duration",
	"episode_urls",
	[]string{
		"ALTER TABLE public.episode_urls",
		"ADD duration decimal;",
	},
)
