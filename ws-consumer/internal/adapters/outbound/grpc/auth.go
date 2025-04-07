package grpc

import (
	"context"
	"errors"
	"fmt"

	"real-time-messaging/consumer/config"
	"real-time-messaging/consumer/pkg/logger"

	pb "real-time-messaging/consumer/api/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client pb.AuthServiceClient
	conn   *grpc.ClientConn
	logger *logger.Logger
}

func NewAuth(cfg *config.Config, logger *logger.Logger) (*AuthClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.Grpc.Auth.Host, cfg.Grpc.Auth.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("Failed to connect to auth service", zap.Error(err),
			zap.String("host", cfg.Grpc.Auth.Host),
			zap.Int("port", cfg.Grpc.Auth.Port))
		return nil, err
	}

	logger.Logger.Info("Connected to auth service")

	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *AuthClient) ValidateToken(ctx context.Context, token string) (string, error) {
	resp, err := c.client.ValidateToken(ctx, &pb.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		c.logger.Error("Failed to validate token", zap.Error(err), zap.String("host", c.conn.Target()))
		return "", err
	}

	if resp.UserId == "" {
		err = errors.New("invalid token")
		c.logger.Error("Failed to validate token", zap.Error(err), zap.String("host", c.conn.Target()))
		return "", err
	}

	return resp.UserId, nil
}

func (c *AuthClient) GenerateToken(ctx context.Context, email string) (string, error) {
	resp, err := c.client.GenerateToken(ctx, &pb.GenerateTokenRequest{
		Email: email,
	})
	if err != nil {
		c.logger.Error("Failed to generate token", zap.Error(err), zap.String("host", c.conn.Target()))
		return "", err
	}

	return resp.Token, nil
}

func (c *AuthClient) Close() error {
	c.logger.Info("closing auth client connection")
	err := c.conn.Close()
	if err != nil {
		c.logger.Error("Failed to close connection", zap.Error(err), zap.String("host", c.conn.Target()))
		return err
	}
	return nil
}
