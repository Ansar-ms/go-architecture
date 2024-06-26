package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate time taken
		duration := time.Since(startTime)

		// Get status, endpoint, and method
		status := c.Writer.Status()
		endpoint := c.Request.RequestURI
		method := c.Request.Method

		// Log the details
		log.Printf("Status: %d | Endpoint: %s | Method: %s | Duration: %s", status, endpoint, method, duration)
	}
}

func LogMessage(message string) {
	log.Println(message)
}
