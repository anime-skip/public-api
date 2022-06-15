package validation

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/utils"
)

func SanitizeExternalLinkURL(url string) (string, error) {
	return utils.SanitizeURL(url, utils.SanitizeURLOptions{
		AllowedSchemes:   []string{"https"},
		AllowedHostnames: []string{internal.ExternalServiceAnilist.Hostname()},
	})
}
