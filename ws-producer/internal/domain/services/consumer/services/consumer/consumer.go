package consumer

import (
	"context"
	"real-time-messaging/producer/internal/domain/entities"
	port "real-time-messaging/producer/internal/domain/ports"
	"real-time-messaging/producer/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ConsumerService struct {
	wsPort            port.Websocket
	messageSubscriber port.MessageEventSubscriber
	logger            *logger.Logger
}

func (s *ConsumerService) Consume(c *gin.Context) error {
	s.logger.Info("websocket request received",
		zap.String("connection", c.GetHeader("Connection")),
		zap.String("upgrade", c.GetHeader("Upgrade")),
		zap.String("sec-websocket-key", c.GetHeader("Sec-WebSocket-Key")),
		zap.String("sec-websocket-version", c.GetHeader("Sec-WebSocket-Version")),
	)

	conn, err := s.wsPort.Upgrade(c)
	if err != nil {
		s.logger.Error("failed to upgrade websocket", zap.Error(err))
		return err
	}

	s.wsPort.AddConnection(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.handleWebsocketConnection(ctx, conn, cancel); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *ConsumerService) handleWebsocketConnection(ctx context.Context, conn *websocket.Conn, cancel context.CancelFunc) error {
	messages, err := s.messageSubscriber.SubscribeMessage(ctx)
	if err != nil {
		s.logger.Error("failed to subscribe to message", zap.Error(err))
		return err
	}

	go s.handleClientMessages(conn, cancel)
	go s.handleServerMessages(ctx, conn, messages, cancel)

	return nil
}

func (s *ConsumerService) handleClientMessages(conn *websocket.Conn, cancel context.CancelFunc) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			s.logger.Info("websocket client disconnected", zap.Error(err))
			cancel()
			return
		}
	}
}

func (s *ConsumerService) handleServerMessages(
	ctx context.Context,
	conn *websocket.Conn,
	messages <-chan *entities.Message,
	cancel context.CancelFunc,
) {
	defer func() {
		conn.Close()
		s.logger.Info("websocket connection closed")
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-messages:
			if !ok {
				return
			}

			if err := s.wsPort.Send(conn, msg); err != nil {
				s.logger.Error("failed to send message to websocket connection",
					zap.Error(err),
					zap.String("remote_addr", conn.RemoteAddr().String()))
				cancel()
				return
			}

			s.logger.Info("message sent to websocket connection",
				zap.Any("message", msg),
				zap.String("remote_addr", conn.RemoteAddr().String()))
		}
	}
}
