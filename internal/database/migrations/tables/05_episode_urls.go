package tables

// CreateEpisodeURLsTable inserts the admin user
var CreateEpisodeURLsTable = migrateTable(
	"episode_urls",
	[]string{
		"CREATE TABLE public.episode_urls",
		"(",
		"    url text NOT NULL,",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    created_by_user_id uuid NOT NULL,",
		"    updated_at timestamp with time zone NOT NULL,",
		"    updated_by_user_id uuid NOT NULL,",
		"    episode_id uuid NOT NULL,",
		"    CONSTRAINT episode_urls_pkey PRIMARY KEY (url)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
	},
	[]string{
		"ALTER TABLE public.episode_urls",
		"    OWNER to postgres;",
	},
)
