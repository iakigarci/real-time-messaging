package di

import (
	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/internal/adapters/outbound/nats"
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

// Container holds all the dependencies of the application
type Container struct {
	MessageProducer port.MessageEventPublisher
	WebsocketPort   port.Websocket
	NatsClient      *nats.Client
	Config          *config.Config
	Logger          *logger.Logger
}

func NewContainer(
	cfg *config.Config,
	logger *logger.Logger,
	websocketPort port.Websocket,
	natsClient *nats.Client,
	messageProducer port.MessageEventPublisher,
) *Container {
	return &Container{
		Config:          cfg,
		Logger:          logger,
		WebsocketPort:   websocketPort,
		NatsClient:      natsClient,
		MessageProducer: messageProducer,
	}
}

func (c *Container) Shutdown() error {
	if err := c.NatsClient.Close(); err != nil {
		return err
	}

	return nil
}
