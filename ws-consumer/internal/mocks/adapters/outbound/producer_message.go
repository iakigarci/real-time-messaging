package mocks_adapters_outbound

import (
	"context"
	"real-time-messaging/consumer/internal/domain/events"

	"github.com/stretchr/testify/mock"
)

type MockMessageProducer struct {
	mock.Mock
}

func (m *MockMessageProducer) PublishMessage(ctx context.Context, message events.BaseEvent) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}
