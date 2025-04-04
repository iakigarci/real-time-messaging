package ports

import (
	"context"
	"real-time-messaging/consumer/internal/domain/events"
)

type EventBroker interface {
	Publish(message events.BaseEvent) error
}

type MessageEventPublisher interface {
	PublishMessage(ctx context.Context, message events.BaseEvent) error
}
