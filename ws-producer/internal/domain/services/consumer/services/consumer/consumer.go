package consumer

import (
	"context"
	port "real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConsumerService struct {
	wsPort            port.Websocket
	messageSubscriber port.MessageEventSubscriber
	logger            *logger.Logger
}

func (s *ConsumerService) Consume(c *gin.Context) error {
	s.logger.Info("websocket request received",
		zap.String("connection", c.GetHeader("Connection")),
		zap.String("upgrade", c.GetHeader("Upgrade")),
		zap.String("sec-websocket-key", c.GetHeader("Sec-WebSocket-Key")),
		zap.String("sec-websocket-version", c.GetHeader("Sec-WebSocket-Version")),
	)

	conn, err := s.wsPort.Upgrade(c)
	if err != nil {
		s.logger.Error("failed to upgrade websocket", zap.Error(err))
		return err
	}

	s.wsPort.AddConnection(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	messages, err := s.messageSubscriber.SubscribeMessage(ctx)
	if err != nil {
		s.logger.Error("failed to subscribe to message", zap.Error(err))
		return err
	}

	go func() {
		defer conn.Close()
		for msg := range messages {
			if err := s.wsPort.SendToAll(msg); err != nil {
				s.logger.Error("failed to send message to websocket connections", zap.Error(err))
				return
			}
			s.logger.Info("message published to all websocket connections", zap.Any("message", msg))
		}
	}()

	<-ctx.Done()
	return nil
}
