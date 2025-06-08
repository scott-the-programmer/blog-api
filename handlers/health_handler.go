package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status" example:"healthy"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	Uptime    string `json:"uptime" example:"1h23m45s"`
	Service   string `json:"service" example:"blog-api"`
	Version   string `json:"version" example:"1.0.0"`
}

// ReadinessResponse represents the readiness check response  
type ReadinessResponse struct {
	Status    string `json:"status" example:"ready"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
}

// LivenessResponse represents the liveness check response
type LivenessResponse struct {
	Status    string `json:"status" example:"alive"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
}

// HealthHandler handles health check requests
type HealthHandler struct {
	startTime time.Time
}

// NewHealthHandler creates a new HealthHandler instance
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
	}
}

// HealthCheck returns the health status of the API
// @Summary Health check
// @Description Get the health status of the API including uptime and version
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (hh *HealthHandler) HealthCheck(c *gin.Context) {
	uptime := time.Since(hh.startTime)
	
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"uptime":    uptime.String(),
		"service":   "blog-api",
		"version":   "1.0.0",
	})
}

// ReadinessCheck returns readiness status (can be extended to check dependencies)
// @Summary Readiness check
// @Description Check if the API is ready to serve requests
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} ReadinessResponse
// @Router /health/ready [get]
func (hh *HealthHandler) ReadinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// LivenessCheck returns liveness status (basic ping)
// @Summary Liveness check
// @Description Check if the API is alive and responding
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} LivenessResponse
// @Router /health/live [get]
func (hh *HealthHandler) LivenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
