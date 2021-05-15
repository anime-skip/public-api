package tables

// EpisodeColumnsToStrings converts some fields that were numbers to be strings to hold real world data
var EpisodeColumnsToStrings = migrateTableChange(
	"MODIFY_EPISODES_TABLE__columns_to_strings",
	`
	UPDATE public.episodes SET number=null WHERE number=0;
	UPDATE public.episodes SET absolute_number=null WHERE absolute_number=0;
	UPDATE public.episodes SET season=null WHERE season=0;
	ALTER TABLE public.episodes ALTER COLUMN number SET DATA TYPE text;
	ALTER TABLE public.episodes ALTER COLUMN absolute_number SET DATA TYPE text;
	ALTER TABLE public.episodes ALTER COLUMN season SET DATA TYPE text;
	`,
	`
	ALTER TABLE public.episodes ALTER COLUMN number SET DATA TYPE integer;
	ALTER TABLE public.episodes ALTER COLUMN absolute_number SET DATA TYPE integer;
	ALTER TABLE public.episodes ALTER COLUMN season SET DATA TYPE integer;
	`,
)
