package nats

import (
	"fmt"
	"real-time-messaging/producer/config"
	"real-time-messaging/producer/pkg/logger"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Client struct {
	conn   *nats.Conn
	logger *logger.Logger
}

func NewClient(cfg *config.Config, logger *logger.Logger) (*Client, error) {
	if cfg.NATS.URL == "" {
		return nil, fmt.Errorf("NATS URL not provided")
	}

	conn, err := nats.Connect(cfg.NATS.URL, nats.UserInfo(cfg.NATS.Username, cfg.NATS.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	logger.Info("NATS client initialized", zap.String("url", cfg.NATS.URL))

	return &Client{
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *Client) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}

func (c *Client) Close() error {
	c.conn.Close()
	return nil
}
