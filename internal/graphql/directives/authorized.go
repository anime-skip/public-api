package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
)

func Authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if context, err := utils.GinContext(ctx); err == nil {
		if jwtError, hasJWTError := context.Get(constants.CTX_JWT_ERROR); hasJWTError {
			return nil, fmt.Errorf("%v", jwtError)
		}
	}
	return next(ctx)
}
