package seeders

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var canonTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_CANON)
var mustWatchTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_MUST_WATCH)
var brandingTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_BRANDING)
var introTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_INTRO)
var mixedIntroTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_INTRO_INTRO)
var newIntroTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_new_INTRO)
var recapTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_RECAP)
var fillerTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_FILLER)
var transitionTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_TRANSITION)
var creditsTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_CREDITS)
var mixedCreditsTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_MIXED_CREDITS)
var newCreditsTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_NEW_CREDITS)
var previewTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_PREVIEW)
var titleCardTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_TITLE_CARD)

var timestampTypes = []*entities.TimestampType{
	{
		Name:        "Canon",
		Description: "New plot that has not been revealed",
		BaseEntity:  basicEntity(canonTimestampTypeID),
	},
	{
		Name:        "Must Watch",
		Description: "A non-canon section of the episode that should not be skipped. (I.E. The Chika Dance https://youtu.be/6xKAsZZgskg)",
		BaseEntity:  basicEntity(mustWatchTimestampTypeID),
	},
	{
		Name:        "Branding",
		Description: "The small animation letting you know who made the show",
		BaseEntity:  basicEntity(brandingTimestampTypeID),
	},
	{
		Name:        "Intro",
		Description: "The intro of each episode, generally around 1:30 long",
		BaseEntity:  basicEntity(introTimestampTypeID),
	},
	{
		Name:        "Mixed Intro",
		Description: "The intro at the beginning of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(mixedIntroTimestampTypeID),
	},
	{
		Name:        "New Intro",
		Description: "The first of an intro, sometimes it's nice to watch each of the intros",
		BaseEntity:  basicEntity(newIntroTimestampTypeID),
	},
	{
		Name:        "Recap",
		Description: "A recap of the previous episode",
		BaseEntity:  basicEntity(recapTimestampTypeID),
	},
	{
		Name:        "Filler",
		Description: "Content that has no bearing on the actual story",
		BaseEntity:  basicEntity(fillerTimestampTypeID),
	},
	{
		Name:        "Transition",
		Description: "The small animation show to transition into and out of commercials",
		BaseEntity:  basicEntity(transitionTimestampTypeID),
	},
	{
		Name:        "Credits",
		Description: "The credits/outro at the end of each episode",
		BaseEntity:  basicEntity(creditsTimestampTypeID),
	},
	{
		Name:        "Mixed Credits",
		Description: "The credits/outro at the end of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(mixedCreditsTimestampTypeID),
	},
	{
		Name:        "New Credits",
		Description: "The first of an outro, sometimes it's nice to watch each of the outros",
		BaseEntity:  basicEntity(newCreditsTimestampTypeID),
	},
	{
		Name:        "Preview",
		Description: "The preview for the next episode",
		BaseEntity:  basicEntity(previewTimestampTypeID),
	},
	{
		Name:        "Title Card",
		Description: "A short section of the episode that just displays the name of the episode, where no plot development takes place",
		BaseEntity:  basicEntity(titleCardTimestampTypeID),
	},
}

// SeedTimestampTypes inserts the necessary timestamp types
var SeedTimestampTypes *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_TIMESTAMP_TYPES",
	Migrate: func(db *gorm.DB) error {
		for _, timestampType := range timestampTypes {
			db = db.Save(timestampType)
		}
		return db.Error
	},
	Rollback: func(db *gorm.DB) error {
		for _, timestampType := range timestampTypes {
			db = db.Delete(timestampType)
		}
		return db.Error
	},
}
