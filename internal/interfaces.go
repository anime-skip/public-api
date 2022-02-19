package internal

type Server interface {
	Start() error
}

type Authenticator interface {
	// Authenticate returns AuthenticationDetails or errors out if the token is invalid
	Authenticate(token string) (*AuthenticationDetails, error)
}
