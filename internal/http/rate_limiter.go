package http

import (
	"net/http"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

type rateTracker struct {
	count   uint
	expires time.Time
}

type requestRateLimiter struct {
	rates map[string]*rateTracker
}

func NewRequestRateLimiter() internal.RateLimiter {
	return &requestRateLimiter{
		rates: map[string]*rateTracker{},
	}
}

// GqlMiddleware implements internal.RateLimiter
func (l *requestRateLimiter) GqlMiddleware() graphql.HandlerExtension {
	return nil
}

// HttpMiddleware implements internal.RateLimiter
func (l *requestRateLimiter) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := context.GetAPIClient(r.Context())

		if client.RateLimitRpm != nil {
			existingTracker, ok := l.rates[client.ID]
			now := time.Now()
			if ok && existingTracker.expires.After(now) {
				existingTracker.count += 1
			} else {
				l.rates[client.ID] = &rateTracker{
					count:   1,
					expires: now.Add(1 * time.Minute),
				}
			}
			existingTracker = l.rates[client.ID]
			if existingTracker.count > *client.RateLimitRpm {
				writeGraphqlError(w, "Rate limit exceeded", http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
