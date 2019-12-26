package tables

// CreatePreferencesTable inserts the admin user
var CreatePreferencesTable = migrateTable(
	"preferences",
	[]string{
		"CREATE TABLE public.preferences",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    user_id uuid NOT NULL,",
		"    enable_auto_skip boolean,",
		"    enable_auto_play boolean,",
		"    skip_branding boolean,",
		"    skip_intros boolean,",
		"    skip_new_intros boolean,",
		"    skip_recaps boolean,",
		"    skip_filler boolean,",
		"    skip_canon boolean,",
		"    skip_transitions boolean,",
		"    skip_credits boolean,",
		"    skip_mixed_credits boolean,",
		"    skip_preview boolean,",
		"    skip_title_card boolean,",
		"    CONSTRAINT preferences_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
	},
	[]string{
		"ALTER TABLE public.preferences",
		"    OWNER to postgres;",
	},
)
