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

type EpisodeURLService interface {
}

type EpisodeService interface {
}

type PreferencesService interface {
	GetByUserID(ctx context.Context, UserID uuid.UUID) (Preferences, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newPreferences Preferences) error
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	// GetByUserID(ctx context.Context, UserID uuid.UUID) ([]ShowAdmin, error)
}

type ShowService interface {
}

type TemplateTimestampService interface {
}

type TemplateService interface {
}

type ThirdPartyEpisodeService interface {
}

type ThirdPartyTimestampService interface {
}

type TimestampTypeService interface {
}

type TimestampService interface {
}

type UserService interface {
	GetByID(ctx context.Context, ID uuid.UUID) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	// CreateInTx(ctx context.Context, tx *sqlx.Tx, newUser User) error
	// DeleteInTx(ctx context.Context, tx *sqlx.Tx, user User) (EpisodeURL, error)
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
