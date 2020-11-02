package models

import "time"

type BetterVRVShow struct {
	ID        string    `json:"objectId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	VRVID     string    `json:"seriesId"`
	Title     string    `json:"seriesTitle"`
}
