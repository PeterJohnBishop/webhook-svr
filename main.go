package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"webhook-svr/mail"

	"github.com/gin-gonic/gin"
)

var (
	emailStore   []mail.MessagePayload
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
		var payload mail.MessagePayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		msg, err := json.Marshal(payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error marshaling JSON"})
		}
		fmt.Printf("Email Recieved: %s", msg)

		storeMutex.Lock()
		emailStore = append(emailStore, payload)
		storeMutex.Unlock()

		c.JSON(http.StatusOK, gin.H{"status": "processed"})
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
