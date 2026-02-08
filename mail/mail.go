package mail

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

type ResendEvent struct {
	Type string `json:"type"`
	Data struct {
		EmailID string `json:"email_id"`
	} `json:"data"`
}

func GetEmail(event ResendEvent) (*resend.Email, error) {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RESEND_API_KEY is missing")
	}

	fmt.Printf("DEBUG: Attempting to fetch Email ID: '%s' using Key starting with: %s\n",
		event.Data.EmailID,
		apiKey[0:5]+"...")

	client := resend.NewClient(apiKey)

	// client.Emails.Get returns (*resend.Email, error)
	email, err := client.Emails.Get(event.Data.EmailID)
	if err != nil {
		return nil, err
	}

	return email, nil
}
