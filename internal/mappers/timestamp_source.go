package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ToTimestampSourceInt(value *internal.TimestampSource) int {
	if value != nil {
		switch *value {
		case internal.TimestampSourceBetterVrv:
			return internal.TIMESTAMP_SOURCE_BETTER_VRV
		}
	}
	return internal.TIMESTAMP_SOURCE_ANIME_SKIP
}

func ToTimestampSourceEnum(value int) internal.TimestampSource {
	switch value {
	case internal.TIMESTAMP_SOURCE_BETTER_VRV:
		return internal.TimestampSourceBetterVrv
	}
	return internal.TimestampSourceAnimeSkip
}
