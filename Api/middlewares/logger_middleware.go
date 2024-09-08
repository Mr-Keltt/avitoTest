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

		// Log the incoming request
		shared.Logger.Infof("Incoming request: %s %s from %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())

		// Processing the request
		c.Next()

		// Logging the completion of the request
		duration := time.Since(startTime)
		shared.Logger.Infof("Completed %s %s in %v with status %d", c.Request.Method, c.Request.URL.Path, duration, c.Writer.Status())
	}
}
