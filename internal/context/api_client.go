package context

import (
	"context"

	"anime-skip.com/public-api/internal"
)

var apiClientKey = &contextKey{"api_client"}

func WithAPIClient(ctx context.Context, apiClient internal.APIClient) context.Context {
	return context.WithValue(ctx, apiClientKey, apiClient)
}

func GetAPIClient(ctx context.Context) internal.APIClient {
	return ctx.Value(apiClientKey).(internal.APIClient)
}
