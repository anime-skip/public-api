package tables

// LowercaseAllEmails updates all the stored emails to be lowercase since they should be case insensitive
var LowercaseAllEmails = migrateTableChange(
	"MODIFY_USERS_TABLE__lowercase_emails",
	`
	UPDATE public.users SET email=lower(email);
	`,
	`
	-- Cannot undo this migration
	SELECT count(*) FROM public.users;
	`,
)
