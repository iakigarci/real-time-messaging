package ws

import (
	"net/http"
	"real-time-messaging/consumer/pkg/logger"
	"strconv"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebsocketService struct {
	upgrader websocket.Upgrader
	logger   *logger.Logger
	handlers []func(messageType int, message []byte) error
}

func (s *WebsocketService) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return s.upgrader.Upgrade(w, r, nil)
}

func (s *WebsocketService) ReadMessage(conn *websocket.Conn) {
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				s.logger.Error("unexpected connection close", zap.Error(err))
			}
			break
		}

		for _, handler := range s.handlers {
			if err := handler(messageType, message); err != nil {
				s.logger.Error("error handling message", zap.Error(err))
				continue
			}
			s.logger.InfoAttrs("message handled", map[string]string{
				"message": string(message),
				"type":    strconv.Itoa(messageType),
			})
		}
	}
}
