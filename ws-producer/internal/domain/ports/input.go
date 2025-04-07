package ports

import (
	"real-time-messaging/producer/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Websocket interface {
	Upgrade(c *gin.Context) (*websocket.Conn, error)
	Send(conn *websocket.Conn, message *entities.Message) error
	SendToAll(message *entities.Message) error
	AddConnection(conn *websocket.Conn)
}

type Consumer interface {
	Consume(c *gin.Context) error
}
