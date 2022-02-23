package seeders

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/postgres/migrations/sqlx_migration"
)

var timestampTypes = []*internal.TimestampType{
	{
		Name:        "Canon",
		Description: "New plot that has not been revealed",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_CANON),
	},
	{
		Name:        "Must Watch",
		Description: "A non-canon section of the episode that should not be skipped. (I.E. The Chika Dance https://youtu.be/6xKAsZZgskg)",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_MUST_WATCH),
	},
	{
		Name:        "Branding",
		Description: "The small animation letting you know who made the show",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_BRANDING),
	},
	{
		Name:        "Intro",
		Description: "The intro of each episode, generally around 1:30 long",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_INTRO),
	},
	{
		Name:        "Mixed Intro",
		Description: "The intro at the beginning of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_INTRO_INTRO),
	},
	{
		Name:        "New Intro",
		Description: "The first of an intro, sometimes it's nice to watch each of the intros",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_NEW_INTRO),
	},
	{
		Name:        "Recap",
		Description: "A recap of the previous episode",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_RECAP),
	},
	{
		Name:        "Filler",
		Description: "Content that has no bearing on the actual story",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_FILLER),
	},
	{
		Name:        "Transition",
		Description: "The small animation show to transition into and out of commercials",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_TRANSITION),
	},
	{
		Name:        "Credits",
		Description: "The credits/outro at the end of each episode",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_CREDITS),
	},
	{
		Name:        "Mixed Credits",
		Description: "The credits/outro at the end of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_MIXED_CREDITS),
	},
	{
		Name:        "New Credits",
		Description: "The first of an outro, sometimes it's nice to watch each of the outros",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_NEW_CREDITS),
	},
	{
		Name:        "Preview",
		Description: "The preview for the next episode",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_PREVIEW),
	},
	{
		Name:        "Title Card",
		Description: "A short section of the episode that just displays the name of the episode, where no plot development takes place",
		BaseEntity:  basicEntity(config.TIMESTAMP_ID_TITLE_CARD),
	},
}

// SeedTimestampTypes inserts the necessary timestamp types
var SeedTimestampTypes = &sqlx_migration.Migration{
	ID: "SEED_TIMESTAMP_TYPES",
	Up: func(tx internal.Tx) error {
		for _, timestampType := range timestampTypes {
			err := insertTimestampType(tx, *timestampType)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Down: func(tx internal.Tx) error {
		for _, timestampType := range timestampTypes {
			err := deleteTimestampType(tx, timestampType.ID)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
