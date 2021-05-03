package context_utils

import (
	"context"
	"fmt"

	"anime-skip.com/backend/internal/utils/constants"
)

func UserID(ctx context.Context) (string, error) {
	userID := ctx.Value(constants.CTX_USER_ID)
	if userID == nil {
		return "", fmt.Errorf("500 Internal Error [003]")
	}
	return userID.(string), nil
}
