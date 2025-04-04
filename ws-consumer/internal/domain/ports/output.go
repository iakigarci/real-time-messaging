package ports

import "real-time-messaging/consumer/internal/domain/entities"

type EventBroker interface {
	Publish(message entities.Message) error
}
