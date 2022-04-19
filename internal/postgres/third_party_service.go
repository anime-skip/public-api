package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/mappers"
)

type ThirdPartyService struct {
	db internal.Database
}

func NewThirdPartyService(db internal.Database) internal.ThirdPartyService {
	return &ThirdPartyService{
		db: db,
	}
}

func (s *ThirdPartyService) Name() string {
	return "postgres.ThirdPartyService"
}

func (s *ThirdPartyService) FindEpisodeByName(ctx context.Context, name string) ([]internal.ThirdPartyEpisode, error) {
	episodes, err := getEpisodesByName(ctx, s.db, &name)
	if err != nil {
		return nil, err
	}

	results := []internal.ThirdPartyEpisode{}
	for _, episode := range episodes {
		show, err := getShowByID(ctx, s.db, episode.ShowID)
		if err != nil {
			return nil, err
		}
		timestamps, err := getTimestampsByEpisodeID(ctx, s.db, episode.ID)
		if err != nil {
			return nil, err
		}
		thirdPartyTimestamps := []internal.ThirdPartyTimestamp{}
		for _, timestamp := range timestamps {
			thirdPartyTimestamps = append(thirdPartyTimestamps, mappers.InternalTimestampToThirdPartyTimestamp(timestamp))
		}
		results = append(results, mappers.InternalEpisodeToThirdPartyEpisode(
			episode,
			mappers.InternalShowToThirdPartyShow(show),
			thirdPartyTimestamps,
		))
	}
	return results, err
}
