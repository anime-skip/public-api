package context

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

var directiveServicesKey = &contextKey{"directive_services"}

func WithDirectiveServices(ctx context.Context, services internal.DirectiveServices) context.Context {
	return context.WithValue(ctx, directiveServicesKey, services)
}

func GetDirectiveServices(ctx context.Context) internal.DirectiveServices {
	return ctx.Value(directiveServicesKey).(internal.DirectiveServices)
}
