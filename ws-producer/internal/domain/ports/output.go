package ports

import (
	"context"
	"real-time-messaging/producer/internal/domain/entities"
)

type MessageEventSubscriber interface {
	SubscribeMessage(ctx context.Context) (<-chan *entities.Message, error)
}
