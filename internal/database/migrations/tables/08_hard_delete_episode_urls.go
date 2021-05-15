package tables

// ModifyEpisodeUrlsTableHardDelete makes the episode_urls table a hard delete
var ModifyEpisodeUrlsTableHardDelete = migrateTableChange(
	"MODIFY_EPISODE_URLS_TABLE__hard_delete",
	`
	ALTER TABLE public.episode_urls
		DROP COLUMN deleted_at,
		DROP COLUMN deleted_by_user_id;
	`,
	`
	ALTER TABLE public.episode_urls
		ADD COLUMN deleted_at,
		ADD COLUMN deleted_by_user_id;
	`,
)
