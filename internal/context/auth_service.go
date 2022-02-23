package context

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

var authServiceKey = &contextKey{"auth_service"}

func WithAuthService(ctx context.Context, authenticator internal.AuthService) context.Context {
	return context.WithValue(ctx, authServiceKey, authenticator)
}

func GetAuthService(ctx context.Context) internal.AuthService {
	return ctx.Value(authServiceKey).(internal.AuthService)
}
