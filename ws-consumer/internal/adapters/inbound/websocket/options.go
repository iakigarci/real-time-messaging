package ws

import (
	"real-time-messaging/consumer/pkg/logger"

	"github.com/gorilla/websocket"
)

type Option func(*Websocket)

func NewWebsocket(opts ...Option) *Websocket {
	w := &Websocket{}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

func WithUpgrader(upgrader websocket.Upgrader) Option {
	return func(w *Websocket) {
		w.upgrader = upgrader
	}
}

func WithLogger(logger *logger.Logger) Option {
	return func(w *Websocket) {
		w.logger = logger
	}
}

func WithHandlers(handlers ...func(messageType int, message []byte) error) Option {
	return func(w *Websocket) {
		w.handlers = handlers
	}
}
