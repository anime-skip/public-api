package tables

// CreateTimestampTypesTable inserts the admin user
var CreateTimestampTypesTable = migrateTable(
	"timestamp_types",
	[]string{
		"CREATE TABLE public.timestamp_types",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    created_by_user_id uuid NOT NULL,",
		"    updated_at timestamp with time zone NOT NULL,",
		"    updated_by_user_id uuid NOT NULL,",
		"    deleted_at timestamp with time zone,",
		"    deleted_by_user_id uuid,",
		"    name text,",
		"    description text,",
		"    CONSTRAINT timestamp_types_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"ALTER TABLE public.timestamp_types",
		"    OWNER to postgres;",
	},
)
