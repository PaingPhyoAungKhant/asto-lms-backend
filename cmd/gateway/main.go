package main

import (
	"asto-lms-backend/internal/shared/config"
	"asto-lms-backend/internal/shared/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

type HomeResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	return router
}

func setupGatewayRoutes(r *gin.Engine) {
	r.GET("/health", healthHandler)
	r.GET("/", homeHandler)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := setupRouter()
	setupGatewayRoutes(router)
	cfg := config.Load()

	port := ":" + cfg.Server.Port

	log.Printf("Starting API Gateway on port %s", port)
	log.Printf("Health heck http://localhost%s/health", port)
	log.Printf("Home: http://localhost%s/", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func healthHandler(c *gin.Context) {
	response := HealthResponse{
		Status:    "Healthy",
		Service:   "api-gateway",
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}

func homeHandler(c *gin.Context) {
	response := HomeResponse{
		Message: "Welcome to LMS Academy API Gateway",
		Version: "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}
