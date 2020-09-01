package tables

var LowercaseAllEmails = migrateTableChange(
	"MODIFY_USERS_TABLE__lowercase_emails",
	[]string{
		"UPDATE public.users",
		"SET email=lower(email);",
	},
	[]string{
		// Cannot undo this migration
		"SELECT count(*) FROM public.users;",
	},
)
