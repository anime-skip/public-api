package context_utils

import (
	"context"

	"anime-skip.com/backend/internal/utils/constants"
)

func Role(ctx context.Context) (int, bool) {
	role := ctx.Value(constants.CTX_ROLE)
	if role == nil {
		return -1, false
	}
	return role.(int), true
}
