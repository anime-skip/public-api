package tables

var ModifyEpisodeUrlsTableHardDelete = migrateTableChange(
	"MODIFY_EPISODE_URLS_TABLE__hard_delete",
	[]string{
		"ALTER TABLE public.episode_urls",
		"DROP COLUMN deleted_at,",
		"DROP COLUMN deleted_by_user_id;",
	},
	[]string{
		"ALTER TABLE public.episode_urls",
		"ADD COLUMN deleted_at,",
		"ADD COLUMN deleted_by_user_id;",
	},
)
