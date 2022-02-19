package internal

import (
	"context"
)

type Server interface {
	Start() error
}

type AuthenticationDetails struct {
	IsAdmin  bool
	IsDev    bool
	UserID   string
	ClientId string
}
type Authenticator interface {
	// Authenticate returns AuthenticationDetails or errors out if the token is invalid
	Authenticate(token string) (*AuthenticationDetails, error)
}

type GetUserByIDParams struct {
	UserID string
}
type UserService interface {
	GetUserByID(ctx context.Context, params GetUserByIDParams) (User, error)
}

type Services struct {
	UserService UserService
}
