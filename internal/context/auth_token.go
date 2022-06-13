package context

import (
	"context"
)

var authTokenKey = &contextKey{"auth_token"}

func WithAuthToken(ctx context.Context, authToken string) context.Context {
	return context.WithValue(ctx, authTokenKey, authToken)
}

func GetAuthToken(ctx context.Context) string {
	return ctx.Value(authTokenKey).(string)
}
