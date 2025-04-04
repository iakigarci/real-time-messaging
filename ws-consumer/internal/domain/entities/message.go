package entities

import "time"

type Message struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Type      int       `json:"type"`
}
