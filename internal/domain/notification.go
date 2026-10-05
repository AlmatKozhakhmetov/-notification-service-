package domain

import "time"

type NotificationType string

const (
	TypeEmail    NotificationType = "EMAIL"
	TypeTelegram NotificationType = "TELEGRAM"
)

type Notification struct {
	ID        string           `json:"id"`
	Type      NotificationType `json:"type"`
	Recipient string           `json:"recipient"`
	Message   string           `json:"message"`
	CreatedAt time.Time        `json:"created_at"`
}
