package subscribers

import (
	"context"
	"encoding/json"
	nats_client "real-time-messaging/producer/internal/adapters/outbound/nats"
	"real-time-messaging/producer/internal/domain/entities"
	"real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"

	"go.uber.org/zap"
)

const (
	MessageSubject = string(nats_client.Message) + ".publish"
)

type MessageSubscriber struct {
	subscriber *nats_client.Subscriber
	logger     *logger.Logger
}

type NATSMessage struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Data      struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		Data      struct {
			Content string `json:"content"`
			Type    int    `json:"type"`
		} `json:"data"`
	} `json:"data"`
}

func NewMessageSubscriber(client *nats_client.Client, logger *logger.Logger) ports.MessageEventSubscriber {
	return &MessageSubscriber{
		subscriber: nats_client.NewSubscriber(client),
		logger:     logger,
	}
}

func (s *MessageSubscriber) SubscribeMessage(ctx context.Context) (<-chan *entities.Message, error) {
	natsMessages, err := s.subscriber.Subscribe(ctx, MessageSubject)
	if err != nil {
		return nil, err
	}

	messages := make(chan *entities.Message, 1000)

	go func() {
		defer close(messages)
		for msg := range natsMessages {
			var natsMsg NATSMessage
			if err := json.Unmarshal(msg.Data, &natsMsg); err != nil {
				s.logger.Error("failed to unmarshal message", zap.Error(err))
				continue
			}

			message := &entities.Message{
				Content: natsMsg.Data.Data.Content,
				Type:    natsMsg.Data.Data.Type,
			}
			messages <- message
		}
	}()

	return messages, nil
}
