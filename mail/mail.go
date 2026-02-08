package mail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/resend/resend-go/v2"
)

type ResendEvent struct {
	Type string `json:"type"`
	Data struct {
		EmailID string `json:"email_id"`
	} `json:"data"`
}

type InboundEmailResponse struct {
	Object  string   `json:"object"`
	ID      string   `json:"id"`
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Html    string   `json:"html"`
	Text    string   `json:"text"`
	// Add other fields if needed
}

func GetEmail(event ResendEvent) (*resend.Email, error) {
	apiKey := os.Getenv("RESEND_API_KEY")

	// 1. Manually construct the request
	// Note: If the standard Get() failed, we try to be specific.
	// However, Resend documentation implies /emails/{id} should work for both
	// IF the key has permissions.
	// If the cURL above worked, you can revert to the SDK but MUST ensure
	// the config was updated on Heroku (step 3 below).

	// If the SDK is definitely failing, use this Raw Request:
	url := fmt.Sprintf("https://api.resend.com/emails/%s", event.Data.EmailID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	// 2. Decode into the SDK struct (or our custom one)
	var email resend.Email
	if err := json.NewDecoder(resp.Body).Decode(&email); err != nil {
		return nil, err
	}

	return &email, nil
}
