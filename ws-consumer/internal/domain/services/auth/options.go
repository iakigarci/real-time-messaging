package auth

import (
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

type AuthBuilderOption func(*AuthSvc)

func New(opts ...AuthBuilderOption) port.AuthService {
	options := &AuthSvc{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithAuthService(authService port.Authentication) AuthBuilderOption {
	return func(opts *AuthSvc) {
		opts.authService = authService
	}
}

func WithLogger(logger *logger.Logger) AuthBuilderOption {
	return func(opts *AuthSvc) {
		opts.logger = logger
	}
}
