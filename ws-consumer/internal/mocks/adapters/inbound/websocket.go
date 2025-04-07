package mocks_adapters_inbound

import (
	"real-time-messaging/consumer/internal/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/mock"
)

type MockWebsocket struct {
	mock.Mock
}

func (m *MockWebsocket) Upgrade(c *gin.Context) (*websocket.Conn, error) {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*websocket.Conn), args.Error(1)
}

func (m *MockWebsocket) Receive(conn *websocket.Conn) (*entities.Message, error) {
	args := m.Called(conn)
	return args.Get(0).(*entities.Message), args.Error(1)
}
