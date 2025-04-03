package ws

import (
	"real-time-messaging/consumer/pkg/logger"

	"github.com/gorilla/websocket"
)

type Option func(*WebsocketService)

func NewWebsocketService(opts ...Option) *WebsocketService {
	svc := &WebsocketService{}
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func WithUpgrader(upgrader websocket.Upgrader) Option {
	return func(svc *WebsocketService) {
		svc.upgrader = upgrader
	}
}

func WithLogger(logger *logger.Logger) Option {
	return func(svc *WebsocketService) {
		svc.logger = logger
	}
}

func WithHandlers(handlers ...func(messageType int, message []byte) error) Option {
	return func(svc *WebsocketService) {
		svc.handlers = handlers
	}
}
