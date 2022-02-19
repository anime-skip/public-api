package internal

import (
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Database = *sqlx.DB

type ApiStatus struct {
	Version       string
	Status        string
	Introspection bool
	Playground    bool
}

type GraphQLHandler struct {
	Handler             http.Handler
	EnableIntrospection bool
}

type UUID = string

// BaseEntity defines the common columns that all db structs should hold
type BaseEntity struct {
	ID              UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt       time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID UUID       `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time  `gorm:"not null"`
	UpdatedByUserID UUID       `gorm:"not null;type:uuid"`
	DeletedAt       *time.Time `gorm:""`
	DeletedByUserID *UUID      `gorm:"type:uuid"`
}

type APIClient struct {
	// Custom Soft Delete, not BaseModel
	ID              string     `gorm:"primary_key"`
	CreatedAt       time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID UUID       `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time  `gorm:"not null"`
	UpdatedByUserID UUID       `gorm:"not null;type:uuid"`
	DeletedAt       *time.Time `gorm:""`
	DeletedByUserID *UUID      `gorm:"type:uuid"`

	// Fields
	UserID         UUID
	AppName        string
	Description    string
	AllowedOrigins *string
	RateLimitRPM   *uint
}

type EpisodeURL struct {
	URL             string    `gorm:"primary_key"`
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedByUserID UUID      `gorm:"not null;type:uuid"`
	UpdatedAt       time.Time `gorm:"not null"`
	UpdatedByUserID UUID      `gorm:"not null;type:uuid"`

	Source           int      `gorm:"not null"`
	Duration         *float64 `gorm:"type:decimal"`
	TimestampsOffset *float64 `gorm:"type:decimal"`
	EpisodeID        UUID     `gorm:"not null;type:uuid"`
}

type Episode struct {
	BaseEntity
	Season         *string
	Number         *string
	AbsoluteNumber *string
	Name           *string
	BaseDuration   *float64 `gorm:"type:decimal"`
	ShowID         UUID     `gorm:"not null;type:uuid"`
}

type Preferences struct {
	ID        UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time

	UserID                     UUID `gorm:"not null;type:uuid"`
	EnableAutoSkip             bool
	EnableAutoPlay             bool
	MinimizeToolbarWhenEditing bool
	HideTimelineWhenMinimized  bool
	ColorTheme                 int `gorm:"not null"`

	SkipBranding     bool
	SkipIntros       bool
	SkipNewIntros    bool
	SkipMixedIntros  bool
	SkipRecaps       bool
	SkipFiller       bool
	SkipCanon        bool
	SkipTransitions  bool
	SkipCredits      bool
	SkipNewCredits   bool
	SkipMixedCredits bool
	SkipPreview      bool
	SkipTitleCard    bool
}

type ShowAdmin struct {
	BaseEntity
	ShowID UUID `gorm:"not null;type:uuid"`
	UserID UUID `gorm:"not null;type:uuid"`
}

type Show struct {
	BaseEntity
	Name         string
	OriginalName *string
	Website      *string
	Image        *string
}

type TemplateTimestamp struct {
	TemplateID  UUID `gorm:"not null;type:uuid"`
	TimestampID UUID `gorm:"not null;type:uuid"`
}

type Template struct {
	BaseEntity
	ShowID          UUID `gorm:"not null;type:uuid"`
	Type            int
	Seasons         pq.StringArray `gorm:"type:text[]"`
	SourceEpisodeID UUID           `gorm:"not null;type:uuid"`
}

type TimestampType struct {
	BaseEntity
	Name        string
	Description string
}

type Timestamp struct {
	BaseEntity
	At        float64
	Source    int  `gorm:"not null"`
	TypeID    UUID `gorm:"not null;type:uuid"`
	EpisodeID UUID `gorm:"not null;type:uuid"`
}

type User struct {
	ID            UUID       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:""`
	Username      string
	Email         string
	PasswordHash  string
	ProfileURL    string
	EmailVerified bool
	Role          int
}
