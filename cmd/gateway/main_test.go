package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGatewayHTTP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	setupGatewayRoutes(router)

	t.Run("GET /", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		expectedResponse := HomeResponse{
			Message: "Welcome to LMS Academy API Gateway",
			Version: "1.0.0",
		}

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response HomeResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Message, response.Message)
		assert.Equal(t, expectedResponse.Version, response.Version)
	})

	t.Run("GET /health", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expectedResponse := HealthResponse{
			Status:    "Healthy",
			Service:   "api-gateway",
			Timestamp: time.Now(),
		}
		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response HealthResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Status, response.Status)
		assert.Equal(t, expectedResponse.Service, response.Service)
	})
}
