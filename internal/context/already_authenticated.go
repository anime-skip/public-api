package context

import (
	"context"
)

var alreadyAuthenticatedKey = &contextKey{"already_authenticated"}

func WithAlreadyAuthenticated(ctx context.Context, alreadyAuthenticated bool) context.Context {
	return context.WithValue(ctx, alreadyAuthenticatedKey, alreadyAuthenticated)
}

func GetAlreadyAuthenticated(ctx context.Context) bool {
	value, ok := ctx.Value(alreadyAuthenticatedKey).(bool)
	return ok && value
}
