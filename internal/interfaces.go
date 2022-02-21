package internal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
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

type UserService interface {
	GetByID(ctx context.Context, ID uuid.UUID) (User, error)
	// GetByUsername(ctx context.Context, username string) (User, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newUser User) error
	// DeleteInTx(ctx context.Context, tx *sqlx.Tx, user User) (EpisodeURL, error)
}

type EpisodeURLService interface {
	GetByURL(ctx context.Context, url string) (EpisodeURL, error)
	Create(ctx context.Context, newEpisodeURL EpisodeURL) error
	Update(ctx context.Context, newEpisodeURL EpisodeURL) error
	DeleteInTx(ctx context.Context, tx *sqlx.Tx, episode EpisodeURL) (EpisodeURL, error)
}

type EpisodeService interface {
	GetByID(ctx context.Context, ID uuid.UUID) (Episode, error)
	Create(ctx context.Context, newEpisode Episode) error
	Update(ctx context.Context, newEpisode Episode) error
	DeleteInTx(ctx context.Context, tx *sqlx.Tx, episode Episode) (Episode, error)
}

type PreferencesService interface {
	// GetByID(ctx context.Context, ID uuid.UUID) (Preferences, error)
	GetByUserID(ctx context.Context, UserID uuid.UUID) (Preferences, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newPreferences Preferences) error
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
	// DeleteInTx(ctx context.Context, tx *sqlx.Tx, preferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	GetByID(ctx context.Context, id uuid.UUID) (ShowAdmin, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]ShowAdmin, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]ShowAdmin, error)
	Create(ctx context.Context, newShowAdmin ShowAdmin) error
	Update(ctx context.Context, newShowAdmin ShowAdmin) error
	DeleteInTx(ctx context.Context, tx *sqlx.Tx, episode ShowAdmin) (ShowAdmin, error)
}

func NewServices(
	userService UserService,
	preferencesService PreferencesService,
) Services {
	return Services{
		UserService:        userService,
		PreferencesService: preferencesService,
	}
}

type Services struct {
	UserService        UserService
	PreferencesService PreferencesService
}
