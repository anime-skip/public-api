package context

import (
	"context"
	"errors"

	"anime-skip.com/public-api/internal"
)

var authClaimsKey = &contextKey{"auth_claims"}

func WithAuthClaims(ctx context.Context, claims internal.AuthClaims) context.Context {
	return context.WithValue(ctx, authClaimsKey, claims)
}

func GetAuthClaims(ctx context.Context) (internal.AuthClaims, error) {
	value, ok := ctx.Value(authClaimsKey).(internal.AuthClaims)
	if !ok {
		return internal.AuthClaims{}, errors.New("AuthClaims is not set yet, does this query/mutation use the @authenticated directive?")
	}
	return value, nil
}
