package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputTimestampType(input internal.InputTimestampType, output *internal.TimestampType) {
	output.Name = input.Name
	output.Description = input.Description
}
