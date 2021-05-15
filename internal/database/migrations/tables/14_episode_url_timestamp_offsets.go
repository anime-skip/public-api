package tables

// AddTimestampsOffsetToEpisodeUrls inserts new timestamps_offset column
var AddTimestampsOffsetToEpisodeUrls = migrateTableChange(
	"MODIFY_EPISODE_URLS_TABLE__add_timestamps_offset",
	`
	ALTER TABLE public.episode_urls
		ADD timestamps_offset decimal;
	`,
	`
	ALTER TABLE public.episode_urls
		DROP COLUMN timestamps_offset;
	`,
)
