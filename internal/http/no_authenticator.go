package http

import "anime-skip.com/timestamps-service/internal"

type noAuthenticator struct{}

func NewNoAuthenticator() internal.Authenticator {
	println("Using no authentication...")
	return &noAuthenticator{}
}

func (a *noAuthenticator) Authenticate(token string) (*internal.AuthenticationDetails, error) {
	return &internal.AuthenticationDetails{
		IsAdmin:  true,
		IsDev:    true,
		UserID:   "00000000-0000-0000-000000000000",
		ClientId: "00000000-0000-0000-000000000000",
	}, nil
}
