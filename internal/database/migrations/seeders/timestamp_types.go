package seeders

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var canonTimestampTypeID = uuid.FromStringOrNil("9edc0037-fa4e-47a7-a29a-d9c43368daa8")
var mustWatchTimestampTypeID = uuid.FromStringOrNil("e384759b-3cd2-4824-9569-128363b4452b")
var brandingTimestampTypeID = uuid.FromStringOrNil("97e3629a-95e5-4b1a-9411-73a47c0d0e25")
var introTimestampTypeID = uuid.FromStringOrNil("14550023-2589-46f0-bfb4-152976506b4c")
var mixedIntroTimestampTypeID = uuid.FromStringOrNil("cbb42238-d285-4c88-9e91-feab4bb8ae0a")
var newIntroTimestampTypeID = uuid.FromStringOrNil("679fb610-ff3c-4cf4-83c0-75bcc7fe8778")
var recapTimestampTypeID = uuid.FromStringOrNil("f38ac196-0d49-40a9-8fcf-f3ef2f40f127")
var fillerTimestampTypeID = uuid.FromStringOrNil("c48f1dce-1890-4394-8ce6-c3f5b2f95e5e")
var transitionTimestampTypeID = uuid.FromStringOrNil("9f0c6532-ccae-4238-83ec-a2804fe5f7b0")
var creditsTimestampTypeID = uuid.FromStringOrNil("2a730a51-a601-439b-bc1f-7b94a640ffb9")
var mixedCreditsTimestampTypeID = uuid.FromStringOrNil("6c4ade53-4fee-447f-89e4-3bb29184e87a")
var newCreditsTimestampTypeID = uuid.FromStringOrNil("d839cdb1-21b3-455d-9c21-7ffeb37adbec")
var previewTimestampTypeID = uuid.FromStringOrNil("c7b1eddb-defa-4bc6-a598-f143081cfe4b")
var titleCardTimestampTypeID = uuid.FromStringOrNil("67321535-a4ea-4f21-8bed-fb3c8286b510")

var timestampTypes = []*entities.TimestampType{
	&entities.TimestampType{
		Name:        "Canon",
		Description: "New plot that has not been revealed",
		BaseEntity:  basicEntity(canonTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Must Watch",
		Description: "A non-canon section of the episode that should not be skipped. (I.E. The Chika Dance https://youtu.be/6xKAsZZgskg)",
		BaseEntity:  basicEntity(mustWatchTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Branding",
		Description: "The small animation letting you know who made the show",
		BaseEntity:  basicEntity(brandingTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Intro",
		Description: "The intro of each episode, generally around 1:30 long",
		BaseEntity:  basicEntity(introTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Mixed Intro",
		Description: "The intro at the beginning of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(mixedIntroTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "New Intro",
		Description: "The first of an intro, sometimes it's nice to watch each of the intros",
		BaseEntity:  basicEntity(newIntroTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Recap",
		Description: "A recap of the previous episode",
		BaseEntity:  basicEntity(recapTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Filler",
		Description: "Content that has no bearing on the actual story",
		BaseEntity:  basicEntity(fillerTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Transition",
		Description: "The small animation show to transition into and out of commercials",
		BaseEntity:  basicEntity(transitionTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Credits",
		Description: "The credits/outro at the end of each episode",
		BaseEntity:  basicEntity(creditsTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Mixed Credits",
		Description: "The credits/outro at the end of an episode that is overlaid with plot. Sometimes the last episode of a show does this",
		BaseEntity:  basicEntity(mixedCreditsTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "New Credits",
		Description: "The first of an outro, sometimes it's nice to watch each of the outros",
		BaseEntity:  basicEntity(newCreditsTimestampTypeID),
	},
	&entities.TimestampType{
		Name:        "Preview",
		Description: "The preview for the next episode",
		BaseEntity:  basicEntity(previewTimestampTypeID),
	},
	&entities.TimestampType{
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
