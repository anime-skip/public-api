package server

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

func headerMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("authorization")
	jwt, err := utils.ValidateAuthHeader(authHeader)

	if err != nil {
		c.Set(constants.CTX_JWT_ERROR, err)
	}
	if jwt != nil {
		c.Set(constants.CTX_USER_ID, jwt["userId"])
		c.Set(constants.CTX_ROLE, jwt["role"])
	}

	c.Next()
}

func ginContextToContextMiddleware(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), constants.CTX_GIN_CONTEXT, c)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func corsMiddleware(c *gin.Context) {
	if utils.EnvBool("IS_DEV") {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // TODO - Figure out origins for prod
	}
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}
