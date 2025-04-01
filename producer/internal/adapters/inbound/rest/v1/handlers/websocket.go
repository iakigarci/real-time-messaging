package handlers

import (
	"net/http"

	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	upgrader websocket.Upgrader
	logger   *logger.Logger
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (h *WebsocketHandler) Websocket(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upgrade connection"})
		return
	}
	defer conn.Close()

	// Listen for incoming messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Log error if needed
			}
			break
		}

		// Handle the received message
		// You can process the message here or send it to another service
		if err := conn.WriteMessage(messageType, message); err != nil {
			break
		}
	}
}
