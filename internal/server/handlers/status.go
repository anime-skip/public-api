package handlers

import (
	"net/http"

	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/env"
	"github.com/gin-gonic/gin"
)

// Status is the handler that can be easily reached to tell if the application is running
func Status() gin.HandlerFunc {
	isPlaygroundEnabled := env.ENABLE_PLAYGROUND
	version := constants.VERSION
	if env.IS_DEV || env.IS_STAGED {
		version += constants.VERSION_SUFFIX
	}
	return func(c *gin.Context) {
		statusData := map[string]interface{}{
			"status":        "RUNNING",
			"version":       version,
			"playground":    isPlaygroundEnabled,
			"introspection": env.ENABLE_INTROSPECTION,
		}
		if isPlaygroundEnabled {
			statusData["playgroundPath"] = "/graphiql"
		}
		c.JSON(http.StatusOK, statusData)
	}
}
