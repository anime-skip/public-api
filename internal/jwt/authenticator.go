package jwt

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

type jwtAuthenticator struct{}

func NewJWTAuthenticator() internal.Authenticator {
	log.D("Using Custom JWT Authenticator...")
	return &jwtAuthenticator{}
}

func (a *jwtAuthenticator) Authenticate(token string) (*internal.AuthenticationDetails, error) {
	log.Panic("Not implemented")
	return nil, nil
}
