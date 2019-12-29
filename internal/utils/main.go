package utils

import (
	"context"

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
