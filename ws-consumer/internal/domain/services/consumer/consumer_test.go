package consumer

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/internal/domain/entities"
	apperrors "real-time-messaging/consumer/internal/domain/errors"
	mock_ports "real-time-messaging/consumer/internal/mocks/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

func TestConsumerService_handleWebSocketConnection(t *testing.T) {
	testCases := []struct {
		desc       string
		setupMocks func(*gin.Context, *mock_ports.MockWebsocket)
		expectErr  error
	}{
		{
			desc: "successful websocket upgrade",
			setupMocks: func(c *gin.Context, wsPort *mock_ports.MockWebsocket) {
				mockConn := &websocket.Conn{}
				wsPort.EXPECT().
					Upgrade(gomock.AssignableToTypeOf(&gin.Context{})).
					Return(mockConn, nil)
			},
			expectErr: nil,
		},
		{
			desc: "websocket upgrade fails",
			setupMocks: func(c *gin.Context, wsPort *mock_ports.MockWebsocket) {
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
			eventRepo := mock_ports.NewMockEventRepository(ctrl)

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

			tC.setupMocks(c, wsPort)

			service := &ConsumerService{
				wsPort:          wsPort,
				messageProducer: producer,
				eventRepository: eventRepo,
				logger:          logger,
			}

			conn, err := service.handleWebSocketConnection(c)

			if tC.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tC.expectErr, err)
				assert.Nil(t, conn)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, conn)
			}
		})
	}
}

func TestConsumerService_processMessage(t *testing.T) {
	testCases := []struct {
		desc       string
		message    *entities.Message
		userID     string
		setupMocks func(*mock_ports.MockMessageEventPublisher, *mock_ports.MockEventRepository)
		expectErr  error
	}{
		{
			desc: "successful message processing",
			message: &entities.Message{
				Content: "test message",
			},
			userID: "test-user-id",
			setupMocks: func(producer *mock_ports.MockMessageEventPublisher, eventRepo *mock_ports.MockEventRepository) {
				producer.EXPECT().
					PublishMessage(gomock.Any(), gomock.Any()).
					Return(nil)

				eventRepo.EXPECT().
					CreateEvent(gomock.Any(), gomock.Any(), "test-user-id").
					Return(nil)
			},
			expectErr: nil,
		},
		{
			desc: "failed message publishing",
			message: &entities.Message{
				Content: "test message",
			},
			userID: "test-user-id",
			setupMocks: func(producer *mock_ports.MockMessageEventPublisher, eventRepo *mock_ports.MockEventRepository) {
				producer.EXPECT().
					PublishMessage(gomock.Any(), gomock.Any()).
					Return(errors.New("failed to publish message"))
			},
			expectErr: errors.New("failed to publish message"),
		},
		{
			desc: "failed event creation",
			message: &entities.Message{
				Content: "test message",
			},
			userID: "test-user-id",
			setupMocks: func(producer *mock_ports.MockMessageEventPublisher, eventRepo *mock_ports.MockEventRepository) {
				producer.EXPECT().
					PublishMessage(gomock.Any(), gomock.Any()).
					Return(nil)

				eventRepo.EXPECT().
					CreateEvent(gomock.Any(), gomock.Any(), "test-user-id").
					Return(errors.New("failed to create event"))
			},
			expectErr: errors.New("failed to create event"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			wsPort := mock_ports.NewMockWebsocket(ctrl)
			producer := mock_ports.NewMockMessageEventPublisher(ctrl)
			eventRepo := mock_ports.NewMockEventRepository(ctrl)

			testConfig := &config.Config{
				Logging: config.LogConfig{
					Level:  config.Info,
					Format: "json",
				},
			}
			logger := logger.New(testConfig)

			tC.setupMocks(producer, eventRepo)

			service := &ConsumerService{
				wsPort:          wsPort,
				messageProducer: producer,
				eventRepository: eventRepo,
				logger:          logger,
			}

			err := service.processMessage(tC.message, tC.userID)

			if tC.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tC.expectErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestConsumerService_createEvent(t *testing.T) {
	message := &entities.Message{
		Content: "test message",
	}

	service := &ConsumerService{}

	event := service.createEvent(message)

	assert.NotEmpty(t, event.ID)
	assert.NotEmpty(t, event.CreatedAt)
	assert.Equal(t, message, event.Data)
}
