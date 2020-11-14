package log

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var RequestLogger = gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"%s[ request ] %s%4s%s %s (%s) %d %s%s%s\n",
		dim,
		bold,
		param.Method,
		reset+dim,
		param.Path,
		param.Latency,
		param.StatusCode,
		red,
		param.ErrorMessage,
		reset,
	)
})
