package mail

import (
	"time"
)

type MessagePayload struct {
	CreatedAt time.Time `json:"created_at"`
	Data      EmailData `json:"data"`
	Type      string    `json:"type"`
}

type EmailData struct {
	Attachments []interface{} `json:"attachments"`
	Bcc         []interface{} `json:"bcc"`
	Cc          []interface{} `json:"cc"`
	CreatedAt   time.Time     `json:"created_at"`
	EmailID     string        `json:"email_id"`
	From        string        `json:"from"`
	MessageID   string        `json:"message_id"`
	Subject     string        `json:"subject"`
	To          []string      `json:"to"`
}
