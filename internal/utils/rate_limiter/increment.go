package rate_limiter

import (
	"errors"
	"time"

	"anime-skip.com/backend/internal/database/entities"
)

type rateTracker struct {
	count   uint
	expires time.Time
}

var rates = map[string]*rateTracker{}

func Increment(client entities.APIClient) error {
	if client.RateLimitRPM == nil {
		return nil
	}

	existingTracker, ok := rates[client.ID]
	now := time.Now()
	if ok && existingTracker.expires.After(now) {
		existingTracker.count += 1
	} else {
		rates[client.ID] = &rateTracker{
			count:   1,
			expires: now.Add(1 * time.Minute),
		}
	}
	existingTracker = rates[client.ID]
	if existingTracker.count > *client.RateLimitRPM {
		return errors.New("Rate limit exceeded")
	}
	return nil
}
