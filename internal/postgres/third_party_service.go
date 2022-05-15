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
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.ThirdPartyEpisode, error) {
		episodes, err := findEpisodes(ctx, tx, internal.EpisodesFilter{
			NameContains: &name,
		})
		if err != nil {
			return nil, err
		}

		results := []internal.ThirdPartyEpisode{}
		for _, episode := range episodes {
			show, err := findShow(ctx, tx, internal.ShowsFilter{
				ID: episode.ShowID,
			})
			if err != nil {
				return nil, err
			}
			timestamps, err := findTimestamps(ctx, tx, internal.TimestampsFilter{
				EpisodeID: episode.ID,
			})
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
	})
}
