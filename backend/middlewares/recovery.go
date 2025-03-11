// middlewares/recovery.go
package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recovery middleware recovers from any panics and logs the error
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\n%s", err, debug.Stack())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
