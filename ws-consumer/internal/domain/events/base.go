package events

import "time"

type BaseEvent struct {
	ID        string      `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	Data      interface{} `json:"data"`
}
