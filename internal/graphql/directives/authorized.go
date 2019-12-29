package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
)

func Authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	context := utils.GinContext(ctx)
	if context != nil {
		jwtError, hasJWTError := context.Get(constants.CTX_JWT_ERROR)
		if hasJWTError {
			return nil, fmt.Errorf("%v", jwtError)
		}
	}

	return next(ctx)
}
