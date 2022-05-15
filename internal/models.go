package internal

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type Database = *sql.DB
type Tx = *sql.Tx

type ApiStatus struct {
	Version       string `json:"version"`
	Status        string `json:"status"`
	Introspection bool   `json:"introspection"`
	Playground    bool   `json:"playground"`
}

type GraphQLHandler struct {
	Handler             http.Handler
	EnableIntrospection bool
}

type Pagination struct {
	Offset int
	Limit  int
}

type APIClientsFilter struct {
	Pagination     *Pagination
	ID             *string
	UserID         *uuid.UUID
	IncludeDeleted bool
}

type RecentlyAddedEpisodesFilter struct {
	Pagination
}

type EpisodesFilter struct {
	Pagination     *Pagination
	ID             *uuid.UUID
	Name           *string
	NameContains   *string
	Number         *string
	Season         *string
	AbsoluteNumber *string
	ShowID         *uuid.UUID
	IncludeDeleted bool
	Sort           string
}

type EpisodeURLsFilter struct {
	Pagination *Pagination
	URL        *string
	EpisodeID  *uuid.UUID
}

type PreferencesFilter struct {
	ID     *uuid.UUID
	UserID *uuid.UUID
}

type ShowAdminsFilter struct {
	Pagination     *Pagination
	ID             *uuid.UUID
	ShowID         *uuid.UUID
	UserID         *uuid.UUID
	IncludeDeleted bool
}

type ShowsFilter struct {
	Pagination     *Pagination
	ID             *uuid.UUID
	Name           *string
	NameContains   *string
	Sort           string
	IncludeDeleted bool
}

type TemplatesFilter struct {
	Pagination      *Pagination
	ID              *uuid.UUID
	SourceEpisodeID *uuid.UUID
	ShowID          *uuid.UUID
	Season          *string
	Type            *int
	IncludeDeleted  bool
}

type TemplateTimestampsFilter struct {
	Pagination     *Pagination
	TemplateID     *uuid.UUID
	TimestampID    *uuid.UUID
	IncludeDeleted bool
}

type TimestampsFilter struct {
	Pagination     *Pagination
	ID             *uuid.UUID
	EpisodeID      *uuid.UUID
	IncludeDeleted bool
}

type TimestampTypesFilter struct {
	Pagination     *Pagination
	ID             *uuid.UUID
	IncludeDeleted bool
}

type UsersFilter struct {
	Pagination      *Pagination
	ID              *uuid.UUID
	Username        *string
	Email           *string
	UsernameOrEmail *string
	IncludeDeleted  bool
}

type APIClient struct {
	ID              string
	CreatedAt       time.Time
	CreatedByUserID uuid.UUID
	UpdatedAt       time.Time
	UpdatedByUserID uuid.UUID
	DeletedAt       *time.Time
	DeletedByUserID *uuid.UUID
	UserID          uuid.UUID
	AppName         string
	Description     string
	AllowedOrigins  *string
	RateLimitRPM    *uint
}

type FullUser struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	DeletedAt     *time.Time
	Username      string
	Email         string
	PasswordHash  string
	ProfileURL    string
	EmailVerified bool
	Role          int
}
