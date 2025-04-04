package nats

import (
	"context"
	"encoding/json"
	"real-time-messaging/consumer/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

type ProducerName string

const (
	Message ProducerName = "message"
)

type Producer struct {
	client *Client
}

func NewProducer(client *Client) *Producer {
	return &Producer{
		client: client,
	}
}

func (p *Producer) PublishMessage(ctx context.Context, subject string, message interface{}) error {
	event := events.BaseEvent{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Data:      message,
	}
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.client.Publish(subject, value)
}

func (p *Producer) Close() error {
	return p.client.Close()
}
