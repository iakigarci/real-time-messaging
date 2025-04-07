package handlers

import (
	"net/http"
	port "real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebsocketHandler struct {
	consumerPort port.ConsumerService
	logger       *logger.Logger
}

func NewWebsocketHandler(consumerPort port.ConsumerService, logger *logger.Logger) *WebsocketHandler {
	return &WebsocketHandler{
		consumerPort: consumerPort,
		logger:       logger,
	}
}

// @Summary WebSocket Connection
// @Description This endpoint establishes a WebSocket connection but cannot be tested via Swagger UI. Use a WebSocket client instead
// @Tags websocket
// @Success 101 {string} string "Switching Protocols"
// @Router /v1/ws [get]
// @x-hidden true
func (h *WebsocketHandler) WebsocketReceive(c *gin.Context) {
	if err := h.consumerPort.Consume(c); err != nil {
		h.logger.Error("error consuming websocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
