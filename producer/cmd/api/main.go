package main

import (
	"log"

	"real-time-messaging/producer/config"
	di "real-time-messaging/producer/internal"
	http_gin "real-time-messaging/producer/internal/adapters/inbound/rest"
	httpserver "real-time-messaging/producer/pkg/http"
	"real-time-messaging/producer/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig[config.Config]()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.New(cfg)

	container := getDIContainer(cfg, logger)
	httpServer := startServers(cfg, container)

	if err := <-httpServer.Notify(); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
	}

	shutdown(httpServer, container, logger)
}

func getDIContainer(cfg *config.Config, logger *logger.Logger) *di.Container {
	return di.NewContainer(cfg,
		logger,
	)
}

func startServers(cfg *config.Config, container *di.Container) *httpserver.Server {
	httpServer := http_gin.New(cfg, container)

	server := httpserver.New(cfg, httpServer.Router)
	return server
}

func shutdown(server *httpserver.Server, container *di.Container, log *logger.Logger) {
	if shutdownErr := server.Shutdown(); shutdownErr != nil {
		log.Error("httpServer.Shutdown", zap.Error(shutdownErr))
	}

	if err := container.Shutdown(); err != nil {
		log.Error("container.Shutdown", zap.Error(err))
	}
}
