package cache

import "time"

type TimedCache struct {
	duration  time.Duration
	invalidAt time.Time
	value     interface{}
}

func NewTimedCache(duration time.Duration) *TimedCache {
	return &TimedCache{
		duration:  duration,
		invalidAt: time.Now(),
		value:     nil,
	}
}

func (cache *TimedCache) Get() interface{} {
	now := time.Now()
	if now.After(cache.invalidAt) {
		return nil
	}
	return cache.value
}

func (cache *TimedCache) Set(value interface{}) {
	cache.invalidAt = time.Now().Add(cache.duration)
	cache.value = value
}
