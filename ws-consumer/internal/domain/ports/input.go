package ports

import (
	"context"
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

type UserService interface {
	GetUserByCredentials(ctx context.Context, email, password string) (*entities.User, error)
}

type AuthService interface {
	GenerateToken(ctx context.Context, user *entities.User) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}
