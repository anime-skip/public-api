package internal

import (
	"context"

	"github.com/gofrs/uuid"
)

type Server interface {
	Start() error
}

type AuthenticationDetails struct {
	IsAdmin  bool
	IsDev    bool
	UserID   uuid.UUID
	ClientId uuid.UUID
}
type Authenticator interface {
	// Authenticate returns AuthenticationDetails or errors out if the token is invalid
	Authenticate(token string) (*AuthenticationDetails, error)
}

type APIClientService interface {
}

type GetRecentlyAddedParams struct {
	Pagination
}
type EpisodeService interface {
	GetRecentlyAdded(ctx context.Context, params GetRecentlyAddedParams) ([]Episode, error)
	GetByID(ctx context.Context, id uuid.UUID) (Episode, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]Episode, error)
}

type EpisodeURLService interface {
	GetByURL(ctx context.Context, url string) (EpisodeURL, error)
	GetByEpisodeId(ctx context.Context, episodeID uuid.UUID) ([]EpisodeURL, error)
}

type PreferencesService interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (Preferences, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newPreferences Preferences) error
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	GetByID(ctx context.Context, id uuid.UUID) (ShowAdmin, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]ShowAdmin, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]ShowAdmin, error)
}

type ShowService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Show, error)
}

type TemplateService interface {
}

type TemplateTimestampService interface {
}

type ThirdPartyEpisodeService interface {
}

type ThirdPartyTimestampService interface {
}

type TimestampService interface {
}

type TimestampTypeService interface {
}

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newUser User) error
	// DeleteInTx(ctx context.Context, tx *sqlx.Tx, user User) (EpisodeURL, error)
}

type Services struct {
	APIClientService         APIClientService
	EpisodeService           EpisodeService
	EpisodeURLService        EpisodeURLService
	PreferencesService       PreferencesService
	ShowAdminService         ShowAdminService
	ShowService              ShowService
	TemplateService          TemplateService
	TemplateTimestampService TemplateTimestampService
	TimestampService         TimestampService
	TimestampTypeService     TimestampTypeService
	UserService              UserService
}
