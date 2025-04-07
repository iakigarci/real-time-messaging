package ws

import (
	"errors"
	"real-time-messaging/producer/internal/domain/entities"
	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Websocket struct {
	upgrader    websocket.Upgrader
	logger      *logger.Logger
	handlers    []func(messageType int, message []byte) error
	connections []*websocket.Conn
}

func (w *Websocket) Upgrade(c *gin.Context) (*websocket.Conn, error) {
	if !w.isWebSocketUpgrade(c) {
		return nil, errors.New("not a websocket upgrade request")
	}
	return w.upgrader.Upgrade(c.Writer, c.Request, nil)
}

func (w *Websocket) Send(conn *websocket.Conn, message *entities.Message) error {
	return conn.WriteJSON(message)
}

func (w *Websocket) SendToAll(message *entities.Message) error {
	for _, conn := range w.connections {
		if err := w.Send(conn, message); err != nil {
			return err
		}
	}
	return nil
}

func (w *Websocket) AddConnection(conn *websocket.Conn) {
	w.connections = append(w.connections, conn)
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
