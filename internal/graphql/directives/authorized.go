package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
)

func isAuthorized(ctx context.Context) error {
	fmt.Println("Is authorized?")
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
