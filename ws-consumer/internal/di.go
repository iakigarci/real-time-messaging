package di

import (
	"real-time-messaging/consumer/config"
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

// Container holds all the dependencies of the application
type Container struct {
	WebsocketPort port.Websocket
	Config        *config.Config
	Logger        *logger.Logger
}

func NewContainer(
	cfg *config.Config,
	logger *logger.Logger,
	websocketPort port.Websocket,
) *Container {
	return &Container{
		Config:        cfg,
		Logger:        logger,
		WebsocketPort: websocketPort,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
