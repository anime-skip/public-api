// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type BaseModel interface {
	IsBaseModel()
}

type Episode struct {
	ID              string       `json:"id"`
	CreatedAt       time.Time    `json:"createdAt"`
	CreatedByUserID string       `json:"createdByUserId"`
	CreatedBy       *User        `json:"createdBy"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	UpdatedByUserID string       `json:"updatedByUserId"`
	UpdatedBy       *User        `json:"updatedBy"`
	DeletedAt       *time.Time   `json:"deletedAt"`
	DeletedByUserID *string      `json:"deletedByUserId"`
	DeletedBy       *User        `json:"deletedBy"`
	Season          *int         `json:"season"`
	Number          *int         `json:"number"`
	AbsoluteNumber  *int         `json:"absoluteNumber"`
	Name            *string      `json:"name"`
	Show            *Show        `json:"show"`
	ShowID          string       `json:"showId"`
	Timestamps      []*Timestamp `json:"timestamps"`
}

func (Episode) IsBaseModel() {}

type EpisodeURL struct {
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedByUserID string    `json:"createdByUserId"`
	CreatedBy       *User     `json:"createdBy"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedByUserID string    `json:"updatedByUserId"`
	UpdatedBy       *User     `json:"updatedBy"`
	EpisodeID       string    `json:"episodeId"`
	Episode         *Episode  `json:"episode"`
}

type InputPreferences struct {
	EnableAutoSkip   bool `json:"enableAutoSkip"`
	EnableAutoPlay   bool `json:"enableAutoPlay"`
	SkipBranding     bool `json:"skipBranding"`
	SkipIntros       bool `json:"skipIntros"`
	SkipNewIntros    bool `json:"skipNewIntros"`
	SkipMixedIntros  bool `json:"skipMixedIntros"`
	SkipRecaps       bool `json:"skipRecaps"`
	SkipFiller       bool `json:"skipFiller"`
	SkipCanon        bool `json:"skipCanon"`
	SkipTransitions  bool `json:"skipTransitions"`
	SkipCredits      bool `json:"skipCredits"`
	SkipNewCredits   bool `json:"skipNewCredits"`
	SkipMixedCredits bool `json:"skipMixedCredits"`
	SkipPreview      bool `json:"skipPreview"`
	SkipTitleCard    bool `json:"skipTitleCard"`
}

type InputShow struct {
	Name         string  `json:"name"`
	OriginalName *string `json:"originalName"`
	Website      *string `json:"website"`
	Image        *string `json:"image"`
}

type InputShowAdmin struct {
	ShowID string `json:"showId"`
	UserID string `json:"userId"`
}

type MyUser struct {
	ID            string       `json:"id"`
	CreatedAt     time.Time    `json:"createdAt"`
	DeletedAt     *time.Time   `json:"deletedAt"`
	Username      string       `json:"username"`
	Email         string       `json:"email"`
	ProfileURL    string       `json:"profileUrl"`
	AdminOfShows  []*ShowAdmin `json:"adminOfShows"`
	EmailVerified bool         `json:"emailVerified"`
	Role          Role         `json:"role"`
	Preferences   *Preferences `json:"preferences"`
}

type Preferences struct {
	ID               string     `json:"id"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        time.Time  `json:"updatedAt"`
	DeletedAt        *time.Time `json:"deletedAt"`
	UserID           string     `json:"userId"`
	User             *User      `json:"user"`
	EnableAutoSkip   bool       `json:"enableAutoSkip"`
	EnableAutoPlay   bool       `json:"enableAutoPlay"`
	SkipBranding     bool       `json:"skipBranding"`
	SkipIntros       bool       `json:"skipIntros"`
	SkipNewIntros    bool       `json:"skipNewIntros"`
	SkipMixedIntros  bool       `json:"skipMixedIntros"`
	SkipRecaps       bool       `json:"skipRecaps"`
	SkipFiller       bool       `json:"skipFiller"`
	SkipCanon        bool       `json:"skipCanon"`
	SkipTransitions  bool       `json:"skipTransitions"`
	SkipCredits      bool       `json:"skipCredits"`
	SkipNewCredits   bool       `json:"skipNewCredits"`
	SkipMixedCredits bool       `json:"skipMixedCredits"`
	SkipPreview      bool       `json:"skipPreview"`
	SkipTitleCard    bool       `json:"skipTitleCard"`
}

type Show struct {
	ID              string       `json:"id"`
	CreatedAt       time.Time    `json:"createdAt"`
	CreatedByUserID string       `json:"createdByUserId"`
	CreatedBy       *User        `json:"createdBy"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	UpdatedByUserID string       `json:"updatedByUserId"`
	UpdatedBy       *User        `json:"updatedBy"`
	DeletedAt       *time.Time   `json:"deletedAt"`
	DeletedByUserID *string      `json:"deletedByUserId"`
	DeletedBy       *User        `json:"deletedBy"`
	Name            string       `json:"name"`
	OriginalName    *string      `json:"originalName"`
	Website         *string      `json:"website"`
	Image           *string      `json:"image"`
	Admins          []*ShowAdmin `json:"admins"`
	Episodes        []*Episode   `json:"episodes"`
}

func (Show) IsBaseModel() {}

type ShowAdmin struct {
	ID              string     `json:"id"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedByUserID string     `json:"createdByUserId"`
	CreatedBy       *User      `json:"createdBy"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	UpdatedByUserID string     `json:"updatedByUserId"`
	UpdatedBy       *User      `json:"updatedBy"`
	DeletedAt       *time.Time `json:"deletedAt"`
	DeletedByUserID *string    `json:"deletedByUserId"`
	DeletedBy       *User      `json:"deletedBy"`
	ShowID          string     `json:"showId"`
	Show            *Show      `json:"show"`
	UserID          string     `json:"userId"`
	User            *User      `json:"user"`
}

func (ShowAdmin) IsBaseModel() {}

type Timestamp struct {
	ID              string         `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	CreatedByUserID string         `json:"createdByUserId"`
	CreatedBy       *User          `json:"createdBy"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	UpdatedByUserID string         `json:"updatedByUserId"`
	UpdatedBy       *User          `json:"updatedBy"`
	DeletedAt       *time.Time     `json:"deletedAt"`
	DeletedByUserID *string        `json:"deletedByUserId"`
	DeletedBy       *User          `json:"deletedBy"`
	At              float64        `json:"at"`
	TypeID          string         `json:"typeId"`
	Type            *TimestampType `json:"type"`
	EpisodeID       string         `json:"episodeId"`
	Epiosde         *Episode       `json:"epiosde"`
}

func (Timestamp) IsBaseModel() {}

type TimestampType struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
}

type User struct {
	ID           string       `json:"id"`
	CreatedAt    time.Time    `json:"createdAt"`
	DeletedAt    *time.Time   `json:"deletedAt"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	ProfileURL   string       `json:"profileUrl"`
	AdminOfShows []*ShowAdmin `json:"adminOfShows"`
}

type Role string

const (
	RoleDev   Role = "DEV"
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleDev,
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleDev, RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
