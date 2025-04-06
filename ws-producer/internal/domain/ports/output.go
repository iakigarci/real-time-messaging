package ports

import (
	"context"

	"github.com/nats-io/nats.go"
)

type MessageEventSubscriber interface {
	SubscribeMessage(ctx context.Context) (<-chan *nats.Msg, error)
}
