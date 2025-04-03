package handlers

import (
	port "real-time-messaging/producer/internal/domain/ports"
	httpserver "real-time-messaging/producer/pkg/http"
	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	upgrader         websocket.Upgrader
	websocketService port.WebsocketService
	logger           *logger.Logger
}

func NewWebsocketHandler(upgrader websocket.Upgrader, websocketService port.WebsocketService, logger *logger.Logger) *WebsocketHandler {
	return &WebsocketHandler{
		upgrader:         upgrader,
		websocketService: websocketService,
		logger:           logger,
	}
}

// @Summary Websocket
// @Description Websocket
// @Accept json
// @Produce json
// @Success 201 {object} httpserver.SuccessResponseData
// @Failure 400 {object} httpserver.ErrorResponseData
// @Failure 401 {object} httpserver.ErrorResponseData
// @Router /ws [get]
func (h *WebsocketHandler) WebsocketReceive(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		httpserver.ErrorResponse(c, err)
		return
	}

	h.websocketService.ReadMessage(conn)
	conn.Close()

	httpserver.SuccessResponse(c, "Websocket connected")
}
