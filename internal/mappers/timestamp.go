package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/utils"
)

func ApplyGraphqlInputTimestamp(input internal.InputTimestamp, output *internal.Timestamp) {
	output.At = input.At
	output.TypeID = input.TypeID
	output.Source = utils.ValueOr(input.Source, internal.TimestampSourceAnimeSkip)
}
