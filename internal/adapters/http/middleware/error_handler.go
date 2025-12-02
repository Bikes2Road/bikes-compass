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
		clientIP := c.ClientIP()
		timestamp := time.Now().Format(time.RFC1123)
		statusCode := c.Writer.Status()
		proto := c.Request.Proto
		userAgent := c.Request.UserAgent()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		c.Next()

		latency := time.Since(start)

		log.Printf("%s - [%s] - [%s] %d %s %s %s \"%s\" %s\"\n", clientIP, timestamp, method, statusCode, path, latency, proto, userAgent, errorMessage)
	}
}
