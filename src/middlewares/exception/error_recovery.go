package exception

import (
	"github.com/gin-gonic/gin"
)

// Recovery -> middleware to use custom error response
func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithoutWriter(f)
}

// RecoveryWithoutWriter -> recover panic to custom middleware
func RecoveryWithoutWriter(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				f(c, err)
			}
		}()

		/**
		 * forward to next middleware
		 */
		c.Next()
	}
}
