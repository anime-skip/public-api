package seeders

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var websiteClientID = uuid.FromStringOrNil(constants.TIMESTAMP_ID_TITLE_CARD)

var clientIds = []*entities.APIClient{
	{
		ID:              "th2oogUKrgOf1J8wMSIUPV0IpBMsLOJi",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Anime Skip Web",
		Description:     "https://anime-skip.com",
	},
	{
		ID:              "OB3AfF3fZg9XlZhxtLvhwLhDcevslhnr",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Anime Skip Web Extension",
		Description:     "Chrome: https://apk.rip/6\nFirefox: https://apk.rip/7",
	},
	{
		ID:              "ZGfO0sMF3eCwLYf8yMSCJjlynwNGRXWE",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Shared Client",
		Description:     "Client ID that anyone can use and is listed on the API docs. If you don't want to get client ID, you use this one",
	},
}

// SeedConstantClientIDs inserts the fixed client ids (web, web extension, etc)
var SeedConstantClientIDs *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_CONSTANT_CLIENT_IDS",
	Migrate: func(db *gorm.DB) error {
		for _, clientID := range clientIds {
			db = db.Save(clientID)
		}
		return db.Error
	},
	Rollback: func(db *gorm.DB) error {
		for _, clientID := range clientIds {
			db = db.Delete(clientID)
		}
		return db.Error
	},
}
