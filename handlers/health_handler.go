package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
func (hh *HealthHandler) ReadinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// LivenessCheck returns liveness status (basic ping)
func (hh *HealthHandler) LivenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
