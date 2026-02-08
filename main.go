package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"webhook-svr/mail"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
)

var (
	emailStore   []resend.Email
	payloadStore []string
	storeMutex   sync.Mutex
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":        "listening",
			"mailhook_data": emailStore,
			"webhook_data":  payloadStore,
		})
	})

	// mailhook
	r.POST("/mail", func(c *gin.Context) {
		var event mail.ResendEvent
		var email resend.Email
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		if event.Type != "email.received" {
			c.JSON(http.StatusOK, gin.H{"status": "ignored"})
			return
		}
		email, _ = mail.GetEmail(c, event)

		storeMutex.Lock()
		emailStore = append(emailStore, email)
		storeMutex.Unlock()

		c.JSON(http.StatusOK, gin.H{"status": "processed", "id": email.Id})
	})

	// webhook
	r.POST("/hook", func(c *gin.Context) {
		payload, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read body"})
			return
		}
		fmt.Println("Received webhook:", string(payload))

		c.JSON(http.StatusOK, gin.H{
			"status":    "received",
			"processed": true,
		})
	})
	r.Run(":" + port)
}
