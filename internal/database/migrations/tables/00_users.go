package tables

// CreateUsersTable inserts the admin user
var CreateUsersTable = migrateTable(
	"users",
	[]string{
		"CREATE TABLE public.users",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    deleted_at timestamp with time zone,",
		"    username text,",
		"    email text,",
		"    password_hash text,",
		"    profile_url text,",
		"    email_verified boolean,",
		"    role integer,",
		"    CONSTRAINT users_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"ALTER TABLE public.users",
		"    OWNER to postgres;",
		"",
		"CREATE UNIQUE INDEX \"user_username\"",
		"    ON public.users USING btree",
		"    (\"username\" ASC NULLS LAST);",
	},
)
