package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/internal/domain/entities"
	mock_ports "real-time-messaging/consumer/internal/mocks/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

func TestAuthSvc_GenerateToken(t *testing.T) {
	testCases := []struct {
		desc       string
		user       *entities.User
		setupMocks func(*mock_ports.MockAuthentication)
		expectErr  error
	}{
		{
			desc: "successful token generation",
			user: &entities.User{
				Email: "test@example.com",
			},
			setupMocks: func(auth *mock_ports.MockAuthentication) {
				auth.EXPECT().
					GenerateToken(gomock.Any(), "test@example.com").
					Return("test-token", nil)
			},
			expectErr: nil,
		},
		{
			desc:       "nil user",
			user:       nil,
			setupMocks: func(auth *mock_ports.MockAuthentication) {},
			expectErr:  errors.New("user is nil"),
		},
		{
			desc: "empty email",
			user: &entities.User{
				Email: "",
			},
			setupMocks: func(auth *mock_ports.MockAuthentication) {},
			expectErr:  errors.New("user is nil"),
		},
		{
			desc: "auth service error",
			user: &entities.User{
				Email: "test@example.com",
			},
			setupMocks: func(auth *mock_ports.MockAuthentication) {
				auth.EXPECT().
					GenerateToken(gomock.Any(), "test@example.com").
					Return("", errors.New("auth service error"))
			},
			expectErr: errors.New("auth service error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authService := mock_ports.NewMockAuthentication(ctrl)
			tC.setupMocks(authService)

			testConfig := &config.Config{
				Logging: config.LogConfig{
					Level:  config.Info,
					Format: "json",
				},
			}
			logger := logger.New(testConfig)

			service := &AuthSvc{
				authService: authService,
				logger:      logger,
			}

			token, err := service.GenerateToken(context.Background(), tC.user)

			if tC.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tC.expectErr.Error(), err.Error())
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func TestAuthSvc_ValidateToken(t *testing.T) {
	testCases := []struct {
		desc       string
		token      string
		setupMocks func(*mock_ports.MockAuthentication)
		expectErr  error
	}{
		{
			desc:  "successful token validation",
			token: "valid-token",
			setupMocks: func(auth *mock_ports.MockAuthentication) {
				auth.EXPECT().
					ValidateToken(gomock.Any(), "valid-token").
					Return("user-id", nil)
			},
			expectErr: nil,
		},
		{
			desc:       "empty token",
			token:      "",
			setupMocks: func(auth *mock_ports.MockAuthentication) {},
			expectErr:  errors.New("token is empty"),
		},
		{
			desc:  "auth service error",
			token: "invalid-token",
			setupMocks: func(auth *mock_ports.MockAuthentication) {
				auth.EXPECT().
					ValidateToken(gomock.Any(), "invalid-token").
					Return("", errors.New("auth service error"))
			},
			expectErr: errors.New("auth service error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authService := mock_ports.NewMockAuthentication(ctrl)
			tC.setupMocks(authService)

			testConfig := &config.Config{
				Logging: config.LogConfig{
					Level:  config.Info,
					Format: "json",
				},
			}
			logger := logger.New(testConfig)

			service := &AuthSvc{
				authService: authService,
				logger:      logger,
			}

			userID, err := service.ValidateToken(context.Background(), tC.token)

			if tC.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tC.expectErr.Error(), err.Error())
				assert.Empty(t, userID)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, userID)
			}
		})
	}
}
