package handlers

import (
	"net/http"

	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

// Status is the handler that can be easily reached to tell if the application is running
func Status() gin.HandlerFunc {
	isPlaygroundEnabled := utils.ENV.ENABLE_PLAYGROUND
	return func(c *gin.Context) {
		statusData := map[string]interface{}{
			"status":        "RUNNING",
			"version":       constants.VERSION,
			"playground":    isPlaygroundEnabled,
			"introspection": utils.ENV.ENABLE_INTROSPECTION,
		}
		if isPlaygroundEnabled {
			statusData["playgroundPath"] = "/graphiql"
		}
		c.JSON(http.StatusOK, statusData)
	}
}
