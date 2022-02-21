package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToTimestampSourceInt(value *graphql.TimestampSource) int {
	if value != nil {
		switch *value {
		case graphql.TimestampSourceBetterVrv:
			return internal.TIMESTAMP_SOURCE_BETTER_VRV
		}
	}
	return internal.TIMESTAMP_SOURCE_ANIME_SKIP
}

func ToTimestampSourceEnum(value int) graphql.TimestampSource {
	switch value {
	case internal.TIMESTAMP_SOURCE_BETTER_VRV:
		return graphql.TimestampSourceBetterVrv
	}
	return graphql.TimestampSourceAnimeSkip
}
