package di

import (
	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/pkg/logger"
)

// Container holds all the dependencies of the application
type Container struct {
	Config *config.Config
	Logger *logger.Logger
}

func NewContainer(
	cfg *config.Config,
	logger *logger.Logger,
) *Container {
	return &Container{
		Config: cfg,
		Logger: logger,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
