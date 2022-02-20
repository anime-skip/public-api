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

type GetUserByIDParams struct {
	UserID uuid.UUID
}
type UserService interface {
	GetUserByID(ctx context.Context, params GetUserByIDParams) (User, error)
}

type GetPreferencesByUserIDParams struct {
	UserID uuid.UUID
}
type PreferencesService interface {
	GetPreferencesByUserID(ctx context.Context, params GetPreferencesByUserIDParams) (Preferences, error)
}

type Services struct {
	UserService        UserService
	PreferencesService PreferencesService
}
