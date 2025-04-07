package user

import (
	"context"
	"real-time-messaging/consumer/internal/domain/entities"
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserSvc struct {
	userRepository port.UserRepository
	logger         *logger.Logger
}

func (svc *UserSvc) GetUserByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	user, err := svc.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		svc.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	svc.logger.Info("user", zap.Any("user", user))

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		svc.logger.Error("failed to compare password hash", zap.Error(err))
		return nil, err
	}

	return user, nil
}
