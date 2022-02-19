package tables

var CreateUsersTable = createTable(
	"CREATE_USERS_TABLE",

	"users",

	`CREATE TABLE users (
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at timestamp with time zone,
		
		-- Custom Fields
		username text NOT NULL,
		email text NOT NULL,
		password_hash text NOT NULL,
		profile_url text NOT NULL,
		email_verified boolean NOT NULL,
		role integer NOT NULL,

		-- Constraints
		CONSTRAINT users_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	
	CREATE UNIQUE INDEX "user_username"
	    ON users USING btree
	    ("username" ASC NULLS LAST);`,
)
