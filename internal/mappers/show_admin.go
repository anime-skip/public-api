package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputShowAdmin(input internal.InputShowAdmin, output *internal.ShowAdmin) {
	output.ShowID = input.ShowID
	output.UserID = input.UserID
}
