package handlers

import (
	"net/http"

	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

// Status is the handler that can be easily reached to tell if the application is running
func Status() gin.HandlerFunc {
	isPlaygroundEnabled := utils.EnvBool("ENABLE_PLAYGROUND")
	isIntrospectionEnabled := utils.EnvBool("ENABLE_INTROSPECTION")

	return func(c *gin.Context) {
		statusData := map[string]interface{}{
			"status":        "RUNNING",
			"version":       constants.VERSION,
			"playground":    isPlaygroundEnabled,
			"introspection": isIntrospectionEnabled,
		}
		if isPlaygroundEnabled {
			statusData["playgroundPath"] = "/graphiql"
		}
		c.JSON(http.StatusOK, statusData)
	}
}
