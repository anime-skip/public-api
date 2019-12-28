package tables

// CreateShowAdminsTable inserts the admin user
var CreateShowAdminsTable = migrateTable(
	"show_admins",
	[]string{
		"CREATE TABLE public.show_admins",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    created_by_user_id uuid NOT NULL,",
		"    updated_at timestamp with time zone NOT NULL,",
		"    updated_by_user_id uuid NOT NULL,",
		"    deleted_at timestamp with time zone,",
		"    deleted_by_user_id uuid,",
		"    show_id uuid NOT NULL,",
		"    user_id uuid NOT NULL,",
		"    CONSTRAINT show_admins_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"ALTER TABLE public.show_admins",
		"    OWNER to postgres;",
	},
)
