package context

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

var authenticatorKey = &contextKey{"authenticator"}

func WithAuthenticator(ctx context.Context, authenticator internal.Authenticator) context.Context {
	return context.WithValue(ctx, authenticatorKey, authenticator)
}

func GetAuthenticator(ctx context.Context) internal.Authenticator {
	return ctx.Value(authenticatorKey).(internal.Authenticator)
}
