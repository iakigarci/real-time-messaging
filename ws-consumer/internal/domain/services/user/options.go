package user

import (
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"
)

type UserBuilderOption func(*UserSvc)

func New(opts ...UserBuilderOption) port.UserService {
	options := &UserSvc{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithUserRepository(userRepository port.UserRepository) UserBuilderOption {
	return func(opts *UserSvc) {
		opts.userRepository = userRepository
	}
}

func WithLogger(logger *logger.Logger) UserBuilderOption {
	return func(opts *UserSvc) {
		opts.logger = logger
	}
}
