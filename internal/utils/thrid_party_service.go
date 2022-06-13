package utils

import (
	"context"
	"sync"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

type AggregateThirdPartyService struct {
	services []internal.ThirdPartyService
}

func NewAggregateThirdPartyService(services []internal.ThirdPartyService) internal.ThirdPartyService {
	return &AggregateThirdPartyService{
		services: services,
	}
}

func (s *AggregateThirdPartyService) Name() string {
	return "AggregateThirdPartyService"
}

func (s *AggregateThirdPartyService) FindEpisodeByName(ctx context.Context, name string) ([]internal.ThirdPartyEpisode, error) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(s.services))
	var results []internal.ThirdPartyEpisode

	for _, service := range s.services {
		realService := service
		go func() {
			defer waitGroup.Done()

			r, err := realService.FindEpisodeByName(ctx, name)
			if err != nil {
				log.E("%s failed to find episode by name (%s): %v", realService.Name(), name, err)
			} else {
				results = append(results, r...)
			}
		}()
	}
	waitGroup.Wait()

	return results, nil
}
