package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PrometheusMetrics = gin.WrapH(promhttp.Handler())
