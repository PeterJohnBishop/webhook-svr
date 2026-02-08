package mail

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
)

type ResendEvent struct {
	Type string `json:"type"`
	Data struct {
		EmailID string `json:"email_id"`
	} `json:"data"`
}

func GetEmail(c *gin.Context, event ResendEvent) (resend.Email, bool) {
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	emailPtr, err := client.Emails.Get(event.Data.EmailID)

	if err != nil {
		fmt.Println("Error fetching email:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email content"})
		return resend.Email{}, false
	}

	return *emailPtr, true
}
