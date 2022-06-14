package internal

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

// Custom Aliases

type Database = *sql.DB
type Tx = *sql.Tx

// Custom Models

type ApiStatus struct {
	Version       string `json:"version"`
	Stage         string `json:"stage"`
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
	NameContains   *string
	Sort           string
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

type ExternalLinksFilter struct {
	URL    *string
	ShowID *uuid.UUID
}

type PreferencesFilter struct {
	ID             *uuid.UUID
	UserID         *uuid.UUID
	IncludeDeleted bool
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
	Type            *TemplateType
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

type FullUser struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	DeletedAt     *time.Time
	Username      string
	Email         string
	PasswordHash  string
	ProfileURL    string
	EmailVerified bool
	Role          Role
}

// Generated Model Implementations

// Value implements the driver.Valuer interface.
func (episodeSource *EpisodeSource) Value() (driver.Value, error) {
	if episodeSource == nil {
		return nil, nil
	}
	switch *episodeSource {
	case EpisodeSourceFunimation:
		return driver.Value(EPISODE_SOURCE_FUNIMATION), nil
	case EpisodeSourceVrv:
		return driver.Value(EPISODE_SOURCE_VRV), nil
	case EpisodeSourceCrunchyroll:
		return driver.Value(EPISODE_SOURCE_CRUNCHYROLL), nil
	default:
		return driver.Value(EPISODE_SOURCE_UNKNOWN), nil
	}
}

// Scan implements the Scanner interface in database/sql package.
func (episodeSource *EpisodeSource) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i, ok := src.(int64)
	if !ok {
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning EpisodeSource value: %T", src),
			Op:      "*EpisodeSource.Scan",
		}
	}
	switch i {
	case EPISODE_SOURCE_FUNIMATION:
		*episodeSource = EpisodeSourceFunimation
	case EPISODE_SOURCE_VRV:
		*episodeSource = EpisodeSourceVrv
	case EPISODE_SOURCE_CRUNCHYROLL:
		*episodeSource = EpisodeSourceCrunchyroll
	default:
		*episodeSource = EpisodeSourceUnknown
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (colorTheme *ColorTheme) Value() (driver.Value, error) {
	if colorTheme == nil {
		return nil, nil
	}
	switch *colorTheme {
	case ColorThemePerService:
		return driver.Value(THEME_PER_SERVICE), nil
	case ColorThemeAnimeSkipBlue:
		return driver.Value(THEME_ANIME_SKIP_BLUE), nil
	case ColorThemeVrvYellow:
		return driver.Value(THEME_VRV_YELLOW), nil
	case ColorThemeFunimationPurple:
		return driver.Value(THEME_FUNIMATION_PURPLE), nil
	case ColorThemeCrunchyrollOrange:
		return driver.Value(THEME_CRUNCHYROLL_ORANGE), nil
	default:
		return nil, &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Unknown ColorTheme: %v", *colorTheme),
			Op:      "*ColorTheme.Value",
		}
	}
}

// Scan implements the Scanner interface in database/sql package.
func (colorTheme *ColorTheme) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i, ok := src.(int64)
	if !ok {
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning ColorTheme value: %T", src),
			Op:      "*ColorTheme.Scan",
		}
	}
	switch i {
	case THEME_PER_SERVICE:
		*colorTheme = ColorThemePerService
	case THEME_ANIME_SKIP_BLUE:
		*colorTheme = ColorThemeAnimeSkipBlue
	case THEME_VRV_YELLOW:
		*colorTheme = ColorThemeVrvYellow
	case THEME_FUNIMATION_PURPLE:
		*colorTheme = ColorThemeFunimationPurple
	case THEME_CRUNCHYROLL_ORANGE:
		*colorTheme = ColorThemeCrunchyrollOrange
	default:
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning ColorTheme value: %T", src),
			Op:      "*ColorTheme.Scan",
		}
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (timestampSource *TimestampSource) Value() (driver.Value, error) {
	if timestampSource == nil {
		return driver.Value(TIMESTAMP_SOURCE_ANIME_SKIP), nil
	}
	switch *timestampSource {
	case TimestampSourceBetterVrv:
		return driver.Value(TIMESTAMP_SOURCE_BETTER_VRV), nil
	default:
		return driver.Value(TIMESTAMP_SOURCE_ANIME_SKIP), nil
	}
}

// Scan implements the Scanner interface in database/sql package.
func (timestampSource *TimestampSource) Scan(src interface{}) error {
	if src == nil {
		*timestampSource = TimestampSourceAnimeSkip
		return nil
	}
	i, ok := src.(int64)
	if !ok {
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning TimestampSource value: %T", src),
			Op:      "*TimestampSource.Scan",
		}
	}
	switch i {
	case TIMESTAMP_SOURCE_BETTER_VRV:
		*timestampSource = TimestampSourceBetterVrv
	default:
		*timestampSource = TimestampSourceAnimeSkip
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (role *Role) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	switch *role {
	case RoleUser:
		return driver.Value(ROLE_USER), nil
	case RoleAdmin:
		return driver.Value(ROLE_ADMIN), nil
	case RoleDev:
		return driver.Value(ROLE_DEV), nil
	default:
		return nil, &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Unknown Role: %v", *role),
			Op:      "*Role.Value",
		}
	}
}

// Scan implements the Scanner interface in database/sql package.
func (role *Role) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i, ok := src.(int64)
	if !ok {
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning Role value: %T", src),
			Op:      "*Role.Scan",
		}
	}
	switch i {
	case ROLE_USER:
		*role = RoleUser
	case ROLE_ADMIN:
		*role = RoleAdmin
	case ROLE_DEV:
		*role = RoleDev
	default:
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning Role value: %T", src),
			Op:      "*Role.Scan",
		}
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (templateType *TemplateType) Value() (driver.Value, error) {
	if templateType == nil {
		return nil, nil
	}
	switch *templateType {
	case TemplateTypeSeasons:
		return driver.Value(TEMPLATE_TYPE_SEASONS), nil
	case TemplateTypeShow:
		return driver.Value(TEMPLATE_TYPE_SHOW), nil
	default:
		return nil, &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Unknown TemplateType: %v", *templateType),
			Op:      "*TemplateType.Value",
		}
	}
}

// Scan implements the Scanner interface in database/sql package.
func (templateType *TemplateType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i, ok := src.(int64)
	if !ok {
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning TemplateType value: %T", src),
			Op:      "*TemplateType.Scan",
		}
	}
	switch i {
	case TEMPLATE_TYPE_SEASONS:
		*templateType = TemplateTypeSeasons
	case TEMPLATE_TYPE_SHOW:
		*templateType = TemplateTypeShow
	default:
		return &Error{
			Code:    EINTERNAL,
			Message: fmt.Sprintf("Invalid type while scanning TemplateType value: %T", src),
			Op:      "*TemplateType.Scan",
		}
	}
	return nil
}
