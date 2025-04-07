package entities

type Message struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
}
