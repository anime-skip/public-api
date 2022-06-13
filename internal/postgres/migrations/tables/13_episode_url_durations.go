package tables

// AddDurationToEpisodeUrls inserts new duration column
var AddDurationToEpisodeUrls = sqlMigration(
	"MODIFY_EPISODE_URLS_TABLE__add_duration",
	`
	ALTER TABLE episode_urls
		ADD duration decimal;
	`,
	`
	ALTER TABLE episode_urls
		DROP COLUMN duration;
	`,
)
