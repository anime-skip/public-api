package context

import (
	"context"

	"anime-skip.com/public-api/internal"
)

var servicesKey = &contextKey{"injected_services"}

func WithServices(ctx context.Context, services internal.Services) context.Context {
	return context.WithValue(ctx, servicesKey, services)
}

func GetServices(ctx context.Context) internal.Services {
	return ctx.Value(servicesKey).(internal.Services)
}
