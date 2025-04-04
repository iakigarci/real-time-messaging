package ws

import (
	"errors"
	"net/http"
	"real-time-messaging/consumer/internal/domain/entities"
	"real-time-messaging/consumer/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Websocket struct {
	upgrader websocket.Upgrader
	logger   *logger.Logger
	handlers []func(messageType int, message []byte) error
}

func (w *Websocket) Upgrade(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	return w.upgrader.Upgrade(writer, request, nil)
}

func (w *Websocket) Read(conn *websocket.Conn) (entities.Message, error) {
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				w.logger.Error("unexpected connection close", zap.Error(err))
			}
			break
		}

		if len(w.handlers) == 0 {
			w.logger.Warn("no message handlers registered")
			continue
		}

		for _, handler := range w.handlers {
			if err := handler(messageType, message); err != nil {
				w.logger.Error("error handling message", zap.Error(err))
				continue
			}
			w.logger.Info("message handled",
				zap.String("message", string(message)),
				zap.Int("type", messageType),
			)
			return entities.Message{
				Content:   string(message),
				Type:      messageType,
				CreatedAt: time.Now(),
			}, nil
		}
	}

	return entities.Message{}, errors.New("no message handlers registered")
}

func (w *Websocket) IsWebSocketUpgrade(c *gin.Context) bool {
	if !websocket.IsWebSocketUpgrade(c.Request) {
		w.logger.Error("not a websocket upgrade request",
			zap.String("connection", c.GetHeader("Connection")),
			zap.String("upgrade", c.GetHeader("Upgrade")),
		)
		return false
	}
	return true
}
