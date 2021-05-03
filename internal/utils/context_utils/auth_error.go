package context_utils

import (
	"context"

	"anime-skip.com/backend/internal/utils/constants"
)

func AuthError(ctx context.Context) error {
	err := ctx.Value(constants.CTX_AUTH_ERROR)
	if err != nil {
		return err.(error)
	}
	return nil
}
