package cache

import "time"

type TimedMapCache struct {
	duration time.Duration
	values   map[string]*TimedCache
}

func NewTimedMapCache(duration time.Duration) *TimedMapCache {
	return &TimedMapCache{
		duration: duration,
		values:   map[string]*TimedCache{},
	}
}

func (cache *TimedMapCache) Get(key string) interface{} {
	if timedCache, ok := cache.values[key]; ok {
		return timedCache.Get()
	}
	return nil
}

func (cache *TimedMapCache) Set(key string, value interface{}) {
	if timedCache, ok := cache.values[key]; ok {
		timedCache.Set(value)
	} else {
		newTimedCache := NewTimedCache(cache.duration)
		newTimedCache.Set(value)
		cache.values[key] = newTimedCache
	}
}
