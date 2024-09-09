package middlewares

import (
	"avitoTest/shared"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLoggerMiddleware - middleware for logging HTTP requests
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Log request details
		shared.Logger.Infof("Incoming request: %s %s from %s, Headers: %v", c.Request.Method, c.Request.URL.Path, c.ClientIP(), c.Request.Header)

		// Processing the request
		c.Next()

		// Log response details
		duration := time.Since(startTime)
		shared.Logger.Infof("Completed %s %s in %v with status %d", c.Request.Method, c.Request.URL.Path, duration, c.Writer.Status())
	}
}
