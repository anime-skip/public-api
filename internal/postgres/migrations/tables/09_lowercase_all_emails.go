package tables

// LowercaseAllEmails updates all the stored emails to be lowercase since they should be case insensitive
var LowercaseAllEmails = sqlMigration(
	"MODIFY_USERS_TABLE__lowercase_emails",
	`
	UPDATE users SET email=lower(email);
	`,
	`
	-- Cannot undo this migration
	SELECT count(*) FROM users;
	`,
)
