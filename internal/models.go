package internal

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Database = *sqlx.DB
type Tx = *sqlx.Tx

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

// BaseEntity defines the common columns that all db structs should hold
type BaseEntity struct {
	ID              uuid.UUID  `                         sql_gen:"primary_key"`
	CreatedAt       time.Time  `db:"created_at"`
	CreatedByUserID uuid.UUID  `db:"created_by_user_id"`
	UpdatedAt       time.Time  `db:"updated_at"`
	UpdatedByUserID uuid.UUID  `db:"updated_by_user_id"`
	DeletedAt       *time.Time `db:"deleted_at"          sql_gen:"soft_delete"`
	DeletedByUserID *uuid.UUID `db:"deleted_by_user_id"`
}

type APIClient struct {
	// Custom Soft Delete, not BaseModel
	ID              string     `                         sql_gen:"primary_key"`
	CreatedAt       time.Time  `db:"created_at"`
	CreatedByUserID uuid.UUID  `db:"created_by_user_id"`
	UpdatedAt       time.Time  `db:"updated_at"`
	UpdatedByUserID uuid.UUID  `db:"updated_by_user_id"`
	DeletedAt       *time.Time `db:"deleted_at"          sql_gen:"soft_delete"`
	DeletedByUserID *uuid.UUID `db:"deleted_by_user_id"`

	// Fields
	UserID         uuid.UUID `db:"user_id"   sql_gen:"get_many"`
	AppName        string    `db:"app_name"`
	Description    string
	AllowedOrigins *string `db:"allowed_origins"`
	RateLimitRPM   *uint   `db:"rate_limit_rpm"`
}

type EpisodeURL struct {
	URL             string    `                         sql_gen:"primary_key"`
	CreatedAt       time.Time `db:"created_at"`
	CreatedByUserID uuid.UUID `db:"created_by_user_id"`
	UpdatedAt       time.Time `db:"updated_at"`
	UpdatedByUserID uuid.UUID `db:"updated_by_user_id"`

	Source           int
	Duration         *float64
	TimestampsOffset *float64  `db:"timestamps_offset"`
	EpisodeID        uuid.UUID `db:"episode_id"         sql_gen:"get_many"`
}

type Episode struct {
	BaseEntity
	Season         *string
	Number         *string
	AbsoluteNumber *string `db:"absolute_number"`
	Name           *string
	BaseDuration   *float64  `db:"base_duration"`
	ShowID         uuid.UUID `db:"show_id"        sql_gen:"get_many"`
}

type Preferences struct {
	ID        uuid.UUID  `                 sql_gen:"primary_key"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"  sql_gen:"soft_delete"`

	UserID                     uuid.UUID `db:"user_id"                        sql_gen:"get_one"`
	EnableAutoSkip             bool      `db:"enable_auto_skip"`
	EnableAutoPlay             bool      `db:"enable_auto_play"`
	MinimizeToolbarWhenEditing bool      `db:"minimize_toolbar_when_editing"`
	HideTimelineWhenMinimized  bool      `db:"hide_timeline_when_minimized"`
	ColorTheme                 int       `db:"color_theme"`

	SkipBranding     bool `db:"skip_branding"`
	SkipIntros       bool `db:"skip_intros"`
	SkipNewIntros    bool `db:"skip_new_intros"`
	SkipMixedIntros  bool `db:"skip_mixed_intros"`
	SkipRecaps       bool `db:"skip_recaps"`
	SkipFiller       bool `db:"skip_filler"`
	SkipCanon        bool `db:"skip_canon"`
	SkipTransitions  bool `db:"skip_transitions"`
	SkipCredits      bool `db:"skip_credits"`
	SkipNewCredits   bool `db:"skip_new_credits"`
	SkipMixedCredits bool `db:"skip_mixed_credits"`
	SkipPreview      bool `db:"skip_preview"`
	SkipTitleCard    bool `db:"skip_title_card"`
}

type ShowAdmin struct {
	BaseEntity
	ShowID uuid.UUID `db:"show_id" sql_gen:"get_many"`
	UserID uuid.UUID `db:"user_id" sql_gen:"get_many"`
}

type Show struct {
	BaseEntity
	Name         string
	OriginalName *string `db:"original_name"`
	Website      *string
	Image        *string
}

type TemplateTimestamp struct {
	TemplateID  uuid.UUID `db:"template_id"  sql_gen:"primary_key,get_many"`
	TimestampID uuid.UUID `db:"timestamp_id" sql_gen:"primary_key,get_one"`
}

type Template struct {
	BaseEntity
	ShowID          uuid.UUID `db:"show_id"           sql_gen:"get_many"`
	Type            int
	Seasons         pq.StringArray
	SourceEpisodeID uuid.UUID `db:"source_episode_id" sql_gen:"get_one"`
}

type TimestampType struct {
	BaseEntity
	Name        string
	Description string
}

type Timestamp struct {
	BaseEntity
	At        float64
	Source    int
	TypeID    uuid.UUID `db:"type_id"`
	EpisodeID uuid.UUID `db:"episode_id" sql_gen:"get_many"`
}

type User struct {
	ID            uuid.UUID  `                     sql_gen:"primary_key"`
	CreatedAt     time.Time  `db:"created_at"`
	DeletedAt     *time.Time `db:"deleted_at"      sql_gen:"soft_delete"`
	Username      string     `                     sql_gen:"get_one"`
	Email         string     `                     sql_gen:"get_one"`
	PasswordHash  string     `db:"password_hash"`
	ProfileURL    string     `db:"profile_url"`
	EmailVerified bool       `db:"email_verified"`
	Role          int
}
