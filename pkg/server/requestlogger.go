package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var customLogger = gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"\x1b[94m[ request ]\x1b[0m \x1b[1m%4s\x1b[0m %s (%s)\x1b[91m%s\x1b[0m\n",
		param.Method,
		param.Path,
		param.Latency,
		param.ErrorMessage,
	)
})
