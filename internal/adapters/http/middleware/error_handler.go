package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for centralized error handling and logging
func ErrorHandler() gin.HandlerFunc {
	return gin.Recovery()
}

// Logger is a middleware for logging HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[%s] %s - %d - %v", method, path, statusCode, duration)
	}
}
