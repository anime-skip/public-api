package context

import (
	"context"

	"anime-skip.com/public-api/internal"
)

var authClaimsKey = &contextKey{"auth_claims"}

func WithAuthClaims(ctx context.Context, claims internal.AuthClaims) context.Context {
	return context.WithValue(ctx, authClaimsKey, claims)
}

func GetAuthClaims(ctx context.Context) (internal.AuthClaims, error) {
	value, ok := ctx.Value(authClaimsKey).(internal.AuthClaims)
	if !ok {
		return internal.AuthClaims{}, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "AuthClaims is not set yet, does this query/mutation use the @authenticated or @optionalAuthenticated directive?",
			Op:      "GetAuthClaims",
		}
	}
	return value, nil
}
