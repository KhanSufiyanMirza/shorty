package middlewares

import (
	"fmt"
	"hex/utils/logger"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const LoggingMiddlewareName = "LoggingMiddleware"

func AccessLogger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		requestPath := strings.TrimSuffix(c.Request.RequestURI, "/")
		if strings.HasSuffix(requestPath, "/health") {
			return
		}

		duration := time.Since(start)

		var ginCtxKeys []any
		for k, v := range c.Keys {
			ginCtxKeys = append(ginCtxKeys, k, v)
		}

		fields := append([]any{
			"client_ip", c.ClientIP(),
			"latency", fmt.Sprintf("%v", duration),
			"method", c.Request.Method,
			"path", c.Request.RequestURI,
			"status", fmt.Sprintf("%d", c.Writer.Status()),
			"referrer", c.Request.Referer(),
		}, ginCtxKeys...)

		if c.Writer.Status() >= http.StatusBadRequest &&
			c.Writer.Status() < http.StatusInternalServerError {
			log.Error(LoggingMiddlewareName, "client request error", nil, append(fields, "gin_errors", c.Errors.String())...)
		} else if c.Writer.Status() >= http.StatusInternalServerError {
			log.Error(LoggingMiddlewareName, "internal server error", nil, append(fields, "gin_errors", c.Errors.String())...)
		} else {
			log.Info(LoggingMiddlewareName, "successful", nil, fields...)
		}
	}
}
