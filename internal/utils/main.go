package utils

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

func GinContext(ctx context.Context) *gin.Context {
	ginContext := ctx.Value(constants.CTX_GIN_CONTEXT)
	if ginContext == nil {
		return nil
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil
	}
	return gc
}

func UserIDFromContext(ctx context.Context) (string, error) {
	if context := GinContext(ctx); context != nil {
		if userID, hasUserID := context.Get(constants.CTX_USER_ID); hasUserID {
			return userID.(string), nil
		}
	}
	return "", fmt.Errorf("500 Internal Error [000]")
}
