package context

import (
	"context"
	"errors"

	"anime-skip.com/timestamps-service/internal"
)

var authenticationDetailsKey = &contextKey{"authentication_details"}

func WithAuthenticationDetails(ctx context.Context, details *internal.AuthenticationDetails) context.Context {
	return context.WithValue(ctx, authenticationDetailsKey, details)
}

func GetAuthenticationDetails(ctx context.Context) (*internal.AuthenticationDetails, error) {
	value, ok := ctx.Value(authenticationDetailsKey).(*internal.AuthenticationDetails)
	if !ok {
		return nil, errors.New("AuthenticationDetails is not set yet, does this query/mutation use the @authenticated directive?")
	}
	return value, nil
}
