package consumer

import (
	"context"
	"real-time-messaging/consumer/internal/domain/events"
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConsumerService struct {
	wsPort          port.Websocket
	messageProducer port.MessageEventPublisher
	eventRepository port.EventRepository
	logger          *logger.Logger
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

	go func() {
		for {
			message, err := s.wsPort.Receive(conn)
			if err != nil {
				s.logger.Error("failed to read websocket message", zap.Error(err))
				return
			}
			event := events.BaseEvent{
				ID:        uuid.New().String(),
				CreatedAt: time.Now(),
				Data:      message,
			}

			if err := s.messageProducer.PublishMessage(context.Background(), event); err != nil {
				s.logger.Error("failed to publish message to event broker", zap.Error(err))
				return
			}

			userID, exists := c.Get("user_id")
			if !exists {
				s.logger.Error("user_id not found in context")
				return
			}

			if err := s.eventRepository.CreateEvent(context.Background(), &event, userID.(string)); err != nil {
				s.logger.Error("failed to create event", zap.Error(err))
				return
			}
			s.logger.Info("message published to event broker", zap.Any("message", message))
		}
	}()

	return nil
}
