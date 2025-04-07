package consumer

import (
	port "real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"
)

type Option func(*ConsumerService)

func NewConsumerService(opts ...Option) *ConsumerService {
	c := &ConsumerService{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithLogger(logger *logger.Logger) Option {
	return func(c *ConsumerService) {
		c.logger = logger
	}
}

func WithWebsocket(ws port.Websocket) Option {
	return func(c *ConsumerService) {
		c.wsPort = ws
	}
}

func WithMessageSubscriber(messageSubscriber port.MessageEventSubscriber) Option {
	return func(c *ConsumerService) {
		c.messageSubscriber = messageSubscriber
	}
}
