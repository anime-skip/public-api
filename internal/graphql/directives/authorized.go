package directives

import (
	"context"
	"fmt"

	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/99designs/gqlgen/graphql"
)

func isAuthorized(ctx context.Context) error {
	if context, err := utils.GinContext(ctx); err == nil {
		if jwtError, hasJWTError := context.Get(constants.CTX_JWT_ERROR); hasJWTError {
			return fmt.Errorf("%v", jwtError)
		}
	}
	return nil
}

func Authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := isAuthorized(ctx); err != nil {
		return nil, err
	}
	return next(ctx)
}
