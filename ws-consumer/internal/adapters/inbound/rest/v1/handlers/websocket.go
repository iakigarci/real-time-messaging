package handlers

import (
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"

	"github.com/gin-gonic/gin"
)

type WebsocketHandler struct {
	consumerPort port.Consumer
	logger       *logger.Logger
}

func NewWebsocketHandler(consumerPort port.Consumer, logger *logger.Logger) *WebsocketHandler {
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
	go h.consumerPort.Consume(c)
}
