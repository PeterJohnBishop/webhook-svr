package mail

import (
	"context"
	"time"

	"github.com/resend/resend-go/v3"
)

type MessagePayload struct {
	CreatedAt time.Time `json:"created_at"`
	Data      EmailData `json:"data"`
	Type      string    `json:"type"`
}

type EmailData struct {
	Attachments []Attachment  `json:"attachments"`
	Bcc         []interface{} `json:"bcc"`
	Cc          []interface{} `json:"cc"`
	CreatedAt   time.Time     `json:"created_at"`
	EmailID     string        `json:"email_id"`
	From        string        `json:"from"`
	MessageID   string        `json:"message_id"`
	Subject     string        `json:"subject"`
	To          []string      `json:"to"`
}

type Attachment struct {
	ID                 string `json:"id"`
	Filename           string `json:"filename"`
	ContentType        string `json:"content_type"`
	ContentDisposition string `json:"content_disposition"`
	ContentID          string `json:"content_id"`
}

func GetMail(client *resend.Client, mailID string) (resend.ReceivedEmail, bool) {
	var email resend.ReceivedEmail
	data, err := client.Emails.Receiving.GetWithContext(
		context.TODO(),
		mailID,
	)
	if err != nil {
		return resend.ReceivedEmail{}, false
	}
	email = *data
	return email, true
}
