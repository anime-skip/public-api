package tables

// CreateUsersTable inserts the admin user
var CreateUsersTable = migrateTable(
	"CREATE_USERS_TABLE",
	"users",
	[]string{
		"CREATE TABLE public.users",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    deleted_at timestamp with time zone,",
		"    username text NOT NULL,",
		"    email text NOT NULL,",
		"    password_hash text NOT NULL,",
		"    profile_url text NOT NULL,",
		"    email_verified boolean NOT NULL,",
		"    role integer NOT NULL,",
		"    CONSTRAINT users_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"CREATE UNIQUE INDEX \"user_username\"",
		"    ON public.users USING btree",
		"    (\"username\" ASC NULLS LAST);",
	},
)
