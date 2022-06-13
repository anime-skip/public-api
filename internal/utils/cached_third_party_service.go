package utils

import (
	"context"
	"fmt"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

type CachedThirdPartyEpisodes struct {
	Episodes []internal.ThirdPartyEpisode
	CachedAt time.Time
}

type CachedThirdPartyService struct {
	service  internal.ThirdPartyService
	cache    map[string]CachedThirdPartyEpisodes
	duration time.Duration
}

// var CACHE_DURATION = 30 * time.Minute

func NewCachedThirdPartyService(service internal.ThirdPartyService, duration time.Duration) internal.ThirdPartyService {
	return &CachedThirdPartyService{
		service:  service,
		duration: duration,
		cache:    map[string]CachedThirdPartyEpisodes{},
	}
}

func (s *CachedThirdPartyService) Name() string {
	return fmt.Sprintf("cached(%s)", s.service.Name())
}

func (s *CachedThirdPartyService) FindEpisodeByName(ctx context.Context, name string) ([]internal.ThirdPartyEpisode, error) {
	cachedResult, ok := s.cache[name]
	isCached := ok && cachedResult.CachedAt.Add(s.duration).After(time.Now())
	if isCached {
		log.V("Using cached result from %s for '%s'", s.service.Name(), name)
		return cachedResult.Episodes, nil
	}

	log.V("Fetching latest '%s' from %s", name, s.service.Name())
	start := time.Now()
	remoteResult, err := s.service.FindEpisodeByName(ctx, name)
	log.D("Latest episodes fetched in %s", time.Since(start).String())
	if err != nil || remoteResult == nil {
		return nil, err
	}

	s.cache[name] = CachedThirdPartyEpisodes{
		Episodes: remoteResult,
		CachedAt: time.Now(),
	}
	return remoteResult, nil
}
