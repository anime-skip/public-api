package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/utils"
)

func InternalShowToThirdPartyShow(show internal.Show) internal.ThirdPartyShow {
	return internal.ThirdPartyShow{
		Name:      show.Name,
		UpdatedAt: &show.UpdatedAt,
		CreatedAt: &show.CreatedAt,
	}
}

func InternalTimestampToThirdPartyTimestamp(timestamp internal.Timestamp) internal.ThirdPartyTimestamp {
	return internal.ThirdPartyTimestamp{
		ID:     timestamp.ID,
		At:     timestamp.At,
		TypeID: timestamp.TypeID,
	}
}

func InternalEpisodeToThirdPartyEpisode(episode internal.Episode, show internal.ThirdPartyShow, timestamps []internal.ThirdPartyTimestamp) internal.ThirdPartyEpisode {
	return internal.ThirdPartyEpisode{
		ID:             episode.ID,
		Season:         episode.Season,
		Number:         episode.Number,
		AbsoluteNumber: episode.AbsoluteNumber,
		BaseDuration:   episode.BaseDuration,
		Name:           episode.Name,
		Source:         utils.Ptr(internal.TimestampSourceAnimeSkip),
		ShowID:         episode.ShowID.String(),
		Show:           &show,
		Timestamps:     utils.PtrSlice(timestamps),
	}
}
