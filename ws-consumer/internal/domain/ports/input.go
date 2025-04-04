package ports

import (
	"net/http"
	"real-time-messaging/consumer/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Websocket interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	Read(conn *websocket.Conn) (entities.Message, error)
	IsWebSocketUpgrade(c *gin.Context) bool
}

type Consumer interface {
	Consume(c *gin.Context) error
}
