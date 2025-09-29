// Package middleware
package middleware

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logEntry := map[string]interface{}{
			"timestamp": param.TimeStamp.Format(time.RFC3339),
			"method":    param.Method,
			"path":      param.Path,
			"status":    param.StatusCode,
			"latency":   param.Latency.String(),
			"client_ip": param.ClientIP,
			"error":     param.ErrorMessage,
		}
		jsonData, _ := json.Marshal(logEntry)
		return string(jsonData) + "\n"
	})
}
