package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func IsShowAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	log.Panic("Not implemented")
	return nil, nil
}
