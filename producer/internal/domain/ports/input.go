package ports

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketService interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	ReadMessage(conn *websocket.Conn)
}
