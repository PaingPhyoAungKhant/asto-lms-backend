// Package middleware - Middlewares (cors,logging,auth,etc)
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS - return a gin middleware that handles Cross-origin Resource sharing
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Header", "Origin, Content-Type,Accept,Authorization, X-Requested-With")

		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
