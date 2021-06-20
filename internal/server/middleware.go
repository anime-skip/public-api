package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/auth"
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/env"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/gin-gonic/gin"
)

func headerMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("authorization")
	jwt, err := auth.ValidateAuthHeader(authHeader)

	if err != nil {
		c.Set(constants.CTX_JWT_ERROR, err)
	}
	if jwt != nil {
		c.Set(constants.CTX_USER_ID, jwt["userId"])
		log.V("Set %s to %v", constants.CTX_USER_ID, jwt["userId"])
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
	if env.IS_DEV {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // TODO - Figure out origins for prod
	}
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Client-ID")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}

func banIPMiddleware(c *gin.Context) {
	if utils.StringArrayIncludes(env.BANNED_IP_ADDRESSES, c.ClientIP()) {
		c.Writer.Write([]byte(`{
			"data": null,
			"error": {
				"code": -134,
				"message": "Request failed"
			}
		}`))
		log.E("Request from banned IP: " + c.ClientIP())
		c.Writer.Header().Add("Content-Type", "application/json")
		if env.SLEEP_BAN_IP {
			time.Sleep(20 * time.Second)
		}
		c.AbortWithStatus(http.StatusOK)
	} else {
		c.Next()
	}
}

func loggerMiddleware(c *gin.Context) {
	requestId := c.Request.Header.Get("x-request-id")
	log.V("<<< [request_id=%s] %s %s client_id=%s client_ip=%s", requestId, c.Request.Method, c.Request.URL, c.Request.Header.Get("X-Client-ID"), c.ClientIP())
	start := time.Now()
	if c.Request.URL.Path == "/graphql" && env.LOG_LEVEL <= constants.LOG_LEVEL_VERBOSE {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.W("Failed to read body: %v", err)
		}
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		if err == nil {
			bodyJSON := map[string]interface{}{}
			err = json.Unmarshal(bodyBytes, &bodyJSON)
			if err != nil {
				// Not JSON body
			} else {
				if str, ok := bodyJSON["operationName"].(string); ok {
					log.V("Operation: %s", strings.TrimSpace(str))
				}
				if str, ok := bodyJSON["query"].(string); ok {
					log.V("Query:\n%s", strings.TrimSpace(str))
				}
				if vars, ok := bodyJSON["variables"]; ok {
					if varsMap, isMap := vars.(map[string]interface{}); isMap {
						if _, hasPassword := varsMap["passwordHash"]; hasPassword {
							varsMap["passwordHash"] = "?"
						}
					}
				}
				log.V("Vars: %v", bodyJSON["variables"])
			}
		}
	}
	c.Next()
	log.V(">>> [request_id=%s] status=%s (%v)", requestId, "?", time.Since(start).String())
}
