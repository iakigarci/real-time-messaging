package subscribers

import (
	"context"
	nats_client "real-time-messaging/producer/internal/adapters/outbound/nats"
	"real-time-messaging/producer/internal/domain/ports"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const (
	MessageSubject = string(nats_client.Message) + ".publish"
)

type MessageSubscriber struct {
	subscriber *nats_client.Subscriber
	logger     *zap.Logger
}

func NewMessageSubscriber(client *nats_client.Client, logger *zap.Logger) ports.MessageEventSubscriber {
	return &MessageSubscriber{
		subscriber: nats_client.NewSubscriber(client),
		logger:     logger,
	}
}

func (s *MessageSubscriber) SubscribeMessage(ctx context.Context) (<-chan *nats.Msg, error) {
	return s.subscriber.Subscribe(ctx, MessageSubject)
}
