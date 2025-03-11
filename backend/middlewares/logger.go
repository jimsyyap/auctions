// middlewares/logger.go
package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Log request details
		fmt.Printf("[%s] %s %s | %d | %s | %s\n",
			time.Now().Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			c.ClientIP(),
		)
	}
}
