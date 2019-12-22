package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status is the handler that can be easily reached to tell if the application is running
func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "RUNNING")
	}
}
