package seeders

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/postgres/migrations/sqlx_migration"
)

var rateLimit60 uint = 60

var clients = []internal.APIClient{
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
var SeedKnownClientIDs = &sqlx_migration.Migration{
	ID: "SEED_KNOWN_CLIENT_IDS",
	Up: func(tx internal.Tx) error {
		for _, client := range clients {
			err := insertAPIClient(tx, client)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Down: func(tx internal.Tx) error {
		for _, client := range clients {
			err := insertAPIClient(tx, client)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
