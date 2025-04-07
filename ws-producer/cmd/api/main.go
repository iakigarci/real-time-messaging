package main

import (
	"log"
	"net/http"
	"os"

	"real-time-messaging/producer/config"
	di "real-time-messaging/producer/internal"
	http_gin "real-time-messaging/producer/internal/adapters/inbound/rest"
	ws "real-time-messaging/producer/internal/adapters/inbound/websocket"
	"real-time-messaging/producer/internal/adapters/outbound/nats"
	"real-time-messaging/producer/internal/adapters/outbound/nats/subscribers"
	httpserver "real-time-messaging/producer/pkg/http"
	"real-time-messaging/producer/pkg/logger"

	_ "real-time-messaging/producer/docs"

	"github.com/gorilla/websocket"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"
)

// @title Real-Time Messaging Producer API
// @version 1.0
// @description API for managing real-time messaging
// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	client, err := nats.NewClient(cfg, logger)
	if err != nil {
		logger.Error("Failed to create NATS client", zap.Error(err))
		os.Exit(1)
	}

	messageSubscriber := subscribers.NewMessageSubscriber(client, logger)

	websocketPort := ws.NewWebsocket(
		ws.WithLogger(logger),
		ws.WithUpgrader(websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: true,
		}),
		ws.WithHandlers(func(messageType int, message []byte) error {
			logger.Info("Processing websocket message",
				zap.Int("messageType", messageType),
				zap.ByteString("message", message))

			return nil
		}),
	)

	return di.NewContainer(cfg,
		logger,
		messageSubscriber,
		websocketPort,
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
