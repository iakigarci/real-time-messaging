package di

import (
	"real-time-messaging/producer/config"
	port "real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"
)

// Container holds all the dependencies of the application
type Container struct {
	Config            *config.Config
	Logger            *logger.Logger
	MessageSubscriber port.MessageEventSubscriber
	WebsocketPort     port.Websocket
}

func NewContainer(
	cfg *config.Config,
	logger *logger.Logger,
	messageSubscriber port.MessageEventSubscriber,
	websocketPort port.Websocket,
) *Container {
	return &Container{
		Config:            cfg,
		Logger:            logger,
		MessageSubscriber: messageSubscriber,
		WebsocketPort:     websocketPort,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
