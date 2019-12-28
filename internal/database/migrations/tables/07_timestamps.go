package tables

// CreateTimestampsTable inserts the admin user
var CreateTimestampsTable = migrateTable(
	"timestamps",
	[]string{
		"CREATE TABLE public.timestamps",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    created_by_user_id uuid NOT NULL,",
		"    updated_at timestamp with time zone NOT NULL,",
		"    updated_by_user_id uuid NOT NULL,",
		"    deleted_at timestamp with time zone,",
		"    deleted_by_user_id uuid,",
		"    at numeric,",
		"    type_id uuid NOT NULL,",
		"    episode_id uuid NOT NULL,",
		"    CONSTRAINT timestamps_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"ALTER TABLE public.timestamps",
		"    OWNER to postgres;",
	},
)
