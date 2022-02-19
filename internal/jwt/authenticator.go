package jwt

import (
	"anime-skip.com/timestamps-service/internal"
)

type jwtAuthenticator struct{}

func NewJWTAuthenticator() internal.Authenticator {
	println("Using custom JWT Authenticator...")
	return &jwtAuthenticator{}
}

func (a *jwtAuthenticator) Authenticate(token string) (*internal.AuthenticationDetails, error) {
	panic("Not implemented")
}
