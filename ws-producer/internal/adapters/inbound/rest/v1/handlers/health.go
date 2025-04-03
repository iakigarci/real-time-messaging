package handlers

import (
	httpserver "real-time-messaging/producer/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Description Check if the service is healthy
// @Accept  json
// @Produce  json
// @Success 200 {object} httpserver.SuccessResponseData
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	httpserver.SuccessResponse(c, "healthy")
}
