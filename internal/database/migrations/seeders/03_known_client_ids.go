package seeders

import (
	"anime-skip.com/backend/internal/database/entities"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var rateLimit60 uint = 60

var clients = []*entities.APIClient{
	{
		ID:              "OB3AfF3fZg9XlZhxtLvhwLhDcevslhnr",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Anime Skip Player Web Extension",
		Description:     "Chrome: https://apk.rip/6\nFirefox: https://apk.rip/7",
	},
	{
		ID:              "ZGfO0sMF3eCwLYf8yMSCJjlynwNGRXWE",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Shared Production Client",
		Description:     "Client ID that anyone can use and is listed on the API docs. If you don't want to get client ID, you use this one",
		RateLimitRPM:    &rateLimit60,
	},
	{
		ID:              "th2oogUKrgOf1J8wMSIUPV0IpBMsLOJi",
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		UserID:          adminUUID,
		AppName:         "Anime Skip Web",
		Description:     "Website hosted at https://anime-skip.com",
	},
}

// SeedKnownClientIDs inserts the basic client IDs
var SeedKnownClientIDs *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_KNOWN_CLIENT_IDS",
	Migrate: func(db *gorm.DB) error {
		for _, client := range clients {
			db = db.Save(client)
			if db.Error != nil {
				return db.Error
			}
		}
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		for _, client := range clients {
			db = db.Save(client)
			if db.Error != nil {
				return db.Error
			}
		}
		return nil
	},
}
