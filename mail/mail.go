package mail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func GetEmail(payload MessagePayload) (EmailData, error) {
	apiKey := os.Getenv("RESEND_API_KEY")

	url := fmt.Sprintf("https://api.resend.com/emails/%s", payload.Data.EmailID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return EmailData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return EmailData{}, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var email MessagePayload
	if err := json.NewDecoder(resp.Body).Decode(email); err != nil {
		return EmailData{}, err
	}

	return email.Data, nil
}
