package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyCreateAPIClient(input internal.CreateAPIClient, output *internal.APIClient) {
	output.AppName = input.AppName
	output.Description = input.Description
}
