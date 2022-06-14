package utils

import (
	"fmt"
	"net/url"

	"anime-skip.com/public-api/internal"
)

func SanitizeUrlString(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", &internal.Error{
			Code:    internal.EINVALID,
			Message: "URL is not valid",
			Op:      "SanitizeUrlString",
			Err:     err,
		}
	}
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Hostname(), u.Path), nil
}
