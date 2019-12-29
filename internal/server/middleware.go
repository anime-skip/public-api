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
