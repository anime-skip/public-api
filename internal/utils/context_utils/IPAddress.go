package context_utils

import (
	"context"
	"errors"

	"anime-skip.com/backend/internal/utils/constants"
)

func IPAddress(ctx context.Context) (string, error) {
	ip := ctx.Value(constants.CTX_IP_ADDRESS)
	if ip == nil {
		return "", errors.New("Could not validate IP address")
	}
	return ip.(string), nil
}
