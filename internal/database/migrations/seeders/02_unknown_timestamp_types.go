package seeders

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var unknownTimestampTypeID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_UNKNOWN)

var unknownTimestampType = &entities.TimestampType{
	Name:        "Unknown",
	Description: "A timestamp that was imported from a different system that didn't have enough information to decide what type it was. These are treated as unskippable",
	BaseEntity:  basicEntity(unknownTimestampTypeID),
}

// SeedUnknownTimestampType inserts the a new type, "Unknown"
var SeedUnknownTimestampType *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_UNKNOWN_TIMESTAMP_TYPE",
	Migrate: func(db *gorm.DB) error {
		db = db.Save(unknownTimestampType)
		return db.Error
	},
	Rollback: func(db *gorm.DB) error {
		db = db.Delete(unknownTimestampType)
		return db.Error
	},
}
