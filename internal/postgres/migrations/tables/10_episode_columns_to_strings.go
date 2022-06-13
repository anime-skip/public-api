package tables

// EpisodeColumnsToStrings converts some fields that were numbers to be strings to hold real world data
var EpisodeColumnsToStrings = sqlMigration(
	"MODIFY_EPISODES_TABLE__columns_to_strings",
	`
	UPDATE episodes SET number=null WHERE number=0;
	UPDATE episodes SET absolute_number=null WHERE absolute_number=0;
	UPDATE episodes SET season=null WHERE season=0;
	ALTER TABLE episodes ALTER COLUMN number SET DATA TYPE text;
	ALTER TABLE episodes ALTER COLUMN absolute_number SET DATA TYPE text;
	ALTER TABLE episodes ALTER COLUMN season SET DATA TYPE text;
	`,
	`
	ALTER TABLE episodes ALTER COLUMN number SET DATA TYPE integer;
	ALTER TABLE episodes ALTER COLUMN absolute_number SET DATA TYPE integer;
	ALTER TABLE episodes ALTER COLUMN season SET DATA TYPE integer;
	`,
)
