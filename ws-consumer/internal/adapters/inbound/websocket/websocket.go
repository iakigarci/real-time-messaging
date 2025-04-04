package ws

import (
	"errors"
	"fmt"
	"real-time-messaging/consumer/internal/domain/entities"
	"real-time-messaging/consumer/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Websocket struct {
	upgrader websocket.Upgrader
	logger   *logger.Logger
	handlers []func(messageType int, message []byte) error
}

func (w *Websocket) Upgrade(c *gin.Context) (*websocket.Conn, error) {
	if !w.isWebSocketUpgrade(c) {
		return nil, errors.New("not a websocket upgrade request")
	}
	return w.upgrader.Upgrade(c.Writer, c.Request, nil)
}

func (w *Websocket) Receive(conn *websocket.Conn) (*entities.Message, error) {
	if len(w.handlers) == 0 {
		return nil, fmt.Errorf("no message handlers registered")
	}

	messageType, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	for _, handler := range w.handlers {
		if err := handler(messageType, message); err != nil {
			return nil, err
		}
	}

	return &entities.Message{
		Content: string(message),
		Type:    messageType,
	}, nil
}

func (w *Websocket) isWebSocketUpgrade(c *gin.Context) bool {
	if !websocket.IsWebSocketUpgrade(c.Request) {
		w.logger.Error("not a websocket upgrade request",
			zap.String("connection", c.GetHeader("Connection")),
			zap.String("upgrade", c.GetHeader("Upgrade")),
		)
		return false
	}
	return true
}
