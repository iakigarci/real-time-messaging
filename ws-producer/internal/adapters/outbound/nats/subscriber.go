package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type SubscriberName string

const (
	Message SubscriberName = "message"
)

type Subscriber struct {
	client *Client
}

func NewSubscriber(client *Client) *Subscriber {
	return &Subscriber{
		client: client,
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, subject string) (chan *nats.Msg, error) {
	messages := make(chan *nats.Msg, 1000)

	subscription, err := s.client.conn.ChanSubscribe(subject, messages)
	if err != nil {
		s.client.logger.Error("failed to subscribe to subject", zap.Error(err))
		close(messages)
		return nil, err
	}

	go func() {
		<-ctx.Done()

		s.client.logger.Info("unsubscribing from subject", zap.String("subject", subject))
		subscription.Unsubscribe()
		close(messages)
	}()

	return messages, nil
}

func (s *Subscriber) Close() error {
	return s.client.Close()
}
