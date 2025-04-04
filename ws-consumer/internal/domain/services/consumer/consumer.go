package consumer

import (
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConsumerService struct {
	wsPort          port.Websocket
	eventBrokerPort port.EventBroker
	logger          *logger.Logger
}

func (s *ConsumerService) Consume(c *gin.Context) error {
	s.logger.Info("websocket request received",
		zap.String("connection", c.GetHeader("Connection")),
		zap.String("upgrade", c.GetHeader("Upgrade")),
		zap.String("sec-websocket-key", c.GetHeader("Sec-WebSocket-Key")),
		zap.String("sec-websocket-version", c.GetHeader("Sec-WebSocket-Version")),
	)
	conn, err := s.wsPort.Upgrade(c.Writer, c.Request)
	if err != nil {
		s.logger.Error("failed to upgrade websocket", zap.Error(err))
		return err
	}

	message, err := s.wsPort.Read(conn)
	if err != nil {
		s.logger.Error("failed to read websocket message", zap.Error(err))
		return err
	}

	if err := s.eventBrokerPort.Publish(message); err != nil {
		s.logger.Error("failed to publish message to event broker", zap.Error(err))
		return err
	}

	return nil
}
