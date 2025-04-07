package consumer

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/internal/domain/entities"
	apperrors "real-time-messaging/consumer/internal/domain/errors"
	"real-time-messaging/consumer/internal/domain/events"
	mock_ports "real-time-messaging/consumer/internal/mocks/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

func TestConsumerService_Consume(t *testing.T) {
	testCases := []struct {
		desc       string
		setupMocks func(*gin.Context, *mock_ports.MockWebsocket, *mock_ports.MockMessageEventPublisher)
		expectErr  error
	}{
		{
			desc: "successful websocket connection and message handling",
			setupMocks: func(c *gin.Context, wsPort *mock_ports.MockWebsocket, producer *mock_ports.MockMessageEventPublisher) {
				mockConn := &websocket.Conn{}

				wsPort.EXPECT().
					Upgrade(gomock.AssignableToTypeOf(&gin.Context{})).
					Return(mockConn, nil)

				testMessage := &entities.Message{
					Content: "test message",
				}

				wsPort.EXPECT().
					Receive(mockConn).
					Return(testMessage, nil)

				producer.EXPECT().
					PublishMessage(gomock.Any(), gomock.Any()).
					DoAndReturn(func(_ context.Context, event events.BaseEvent) error {
						assert.NotEmpty(t, event.ID)
						assert.NotEmpty(t, event.CreatedAt)
						assert.Equal(t, testMessage, event.Data)
						return nil
					})

				wsPort.EXPECT().
					Receive(mockConn).
					Return(nil, errors.New("connection closed"))
			},
			expectErr: nil,
		},
		{
			desc: "websocket upgrade fails",
			setupMocks: func(c *gin.Context, wsPort *mock_ports.MockWebsocket, producer *mock_ports.MockMessageEventPublisher) {
				wsPort.EXPECT().
					Upgrade(gomock.AssignableToTypeOf(&gin.Context{})).
					Return(nil, apperrors.ErrNotWebSocketUpgrade)
			},
			expectErr: apperrors.ErrNotWebSocketUpgrade,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			wsPort := mock_ports.NewMockWebsocket(ctrl)
			producer := mock_ports.NewMockMessageEventPublisher(ctrl)

			testConfig := &config.Config{
				Logging: config.LogConfig{
					Level:  config.Info,
					Format: "json",
				},
			}
			logger := logger.New(testConfig)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/ws", nil)
			c.Request.Header.Set("Connection", "Upgrade")
			c.Request.Header.Set("Upgrade", "websocket")
			c.Request.Header.Set("Sec-WebSocket-Key", "test-key")
			c.Request.Header.Set("Sec-WebSocket-Version", "13")

			tC.setupMocks(c, wsPort, producer)

			service := &ConsumerService{
				wsPort:          wsPort,
				messageProducer: producer,
				logger:          logger,
			}

			err := service.Consume(c)

			if tC.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tC.expectErr, err)
			} else {
				assert.NoError(t, err)
				time.Sleep(100 * time.Millisecond)
			}
		})
	}
}
