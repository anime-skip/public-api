package tables

// CreatePreferencesTable inserts the admin user
var CreatePreferencesTable = migrateTable(
	"preferences",
	[]string{
		"CREATE TABLE public.preferences",
		"(",
		"    id uuid NOT NULL DEFAULT uuid_generate_v4(),",
		"    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		"    updated_at timestamp with time zone NOT NULL,",
		"    deleted_at timestamp with time zone,",
		"    user_id uuid NOT NULL,",
		"    enable_auto_skip boolean DEFAULT true,",
		"    enable_auto_play boolean DEFAULT true,",
		"    skip_branding boolean DEFAULT true,",
		"    skip_intros boolean DEFAULT true,",
		"    skip_new_intros boolean DEFAULT false,",
		"    skip_mixed_intros boolean DEFAULT false,",
		"    skip_recaps boolean DEFAULT true,",
		"    skip_filler boolean DEFAULT true,",
		"    skip_canon boolean DEFAULT false,",
		"    skip_transitions boolean DEFAULT true,",
		"    skip_credits boolean DEFAULT true,",
		"    skip_new_credits boolean DEFAULT false,",
		"    skip_mixed_credits boolean DEFAULT false,",
		"    skip_preview boolean DEFAULT true,",
		"    skip_title_card boolean DEFAULT true,",
		"    CONSTRAINT preferences_pkey PRIMARY KEY (id)",
		")",
		"WITH (",
		"    OIDS = FALSE",
		")",
		"TABLESPACE pg_default;",
		"",
		"ALTER TABLE public.preferences",
		"    OWNER to postgres;",
	},
)
