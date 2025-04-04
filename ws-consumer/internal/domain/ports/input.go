package ports

import (
	"real-time-messaging/consumer/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Websocket interface {
	Upgrade(c *gin.Context) (*websocket.Conn, error)
	Receive(conn *websocket.Conn) (*entities.Message, error)
}

type Consumer interface {
	Consume(c *gin.Context) error
}
