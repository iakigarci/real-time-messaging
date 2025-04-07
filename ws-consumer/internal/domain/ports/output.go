package ports

import (
	"context"
	"real-time-messaging/consumer/internal/domain/entities"
	"real-time-messaging/consumer/internal/domain/events"
)

type EventBroker interface {
	Publish(message events.BaseEvent) error
}

type MessageEventPublisher interface {
	PublishMessage(ctx context.Context, message events.BaseEvent) error
}

type Authentication interface {
	ValidateToken(ctx context.Context, token string) (string, error)
	GenerateToken(ctx context.Context, email string) (string, error)
	Close() error
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}
