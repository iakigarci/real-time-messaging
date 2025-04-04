package producers

import (
	"context"
	"real-time-messaging/consumer/internal/adapters/outbound/nats"
	"real-time-messaging/consumer/internal/domain/events"
	port "real-time-messaging/consumer/internal/domain/ports"
)

type MessageProducer struct {
	producer *nats.Producer
}

func NewMessageProducer(client *nats.Client) port.MessageEventPublisher {
	return &MessageProducer{
		producer: nats.NewProducer(client),
	}
}

func (p *MessageProducer) PublishMessage(ctx context.Context, message events.BaseEvent) error {
	return p.producer.PublishMessage(ctx, string(nats.Message)+".publish", message)
}
