package errors

import "errors"

var (
	ErrNotWebSocketUpgrade         = errors.New("not a websocket upgrade request")
	ErrNoMessageHandlersRegistered = errors.New("no message handlers registered")
)
