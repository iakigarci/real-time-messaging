package auth

import (
	"context"
	"errors"
	"real-time-messaging/consumer/internal/domain/entities"
	port "real-time-messaging/consumer/internal/domain/ports"
	"real-time-messaging/consumer/pkg/logger"

	"go.uber.org/zap"
)

type AuthSvc struct {
	authService port.Authentication
	logger      *logger.Logger
}

func (svc *AuthSvc) GenerateToken(ctx context.Context, user *entities.User) (string, error) {
	if user == nil || user.Email == "" {
		return "", errors.New("user is nil")
	}

	token, err := svc.authService.GenerateToken(ctx, user.Email)
	if err != nil {
		svc.logger.Error("failed to generate token", zap.Error(err))
		return "", err
	}

	return token, nil
}

func (svc *AuthSvc) ValidateToken(ctx context.Context, token string) (string, error) {
	if token == "" {
		return "", errors.New("token is empty")
	}

	userID, err := svc.authService.ValidateToken(ctx, token)
	if err != nil {
		svc.logger.Error("failed to validate token", zap.Error(err))
		return "", err
	}

	return userID, nil
}
