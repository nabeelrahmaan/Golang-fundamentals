package middlewear

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func LoginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		fmt.Printf("[Login-Log] %s %s | Status: %d | Duration: %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}