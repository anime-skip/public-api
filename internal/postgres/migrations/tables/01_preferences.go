package tables

var CreatePreferencesTable = createTable(
	"CREATE_PREFERENCES_TABLE",

	"preferences",

	`CREATE TABLE preferences (
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp with time zone NOT NULL,
		deleted_at timestamp with time zone,
		
		-- Custom Fields
		user_id uuid NOT NULL,
		enable_auto_skip boolean,
		enable_auto_play boolean,
		skip_branding boolean,
		skip_intros boolean,
		skip_new_intros boolean,
		skip_mixed_intros boolean,
		skip_recaps boolean,
		skip_filler boolean,
		skip_canon boolean,
		skip_transitions boolean,
		skip_credits boolean,
		skip_new_credits boolean,
		skip_mixed_credits boolean,
		skip_preview boolean,
		skip_title_card boolean,

		-- Constraints
		CONSTRAINT preferences_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
