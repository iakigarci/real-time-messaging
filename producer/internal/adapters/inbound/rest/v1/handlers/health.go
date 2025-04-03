package handlers

import (
	httpserver "real-time-messaging/producer/pkg/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	httpserver.SuccessResponse(c, "healthy")
}
