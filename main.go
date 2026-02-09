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

	r.POST("/webhook", func(c *gin.Context) {

		bodyBytes, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read body"})
			return
		}

		var mailPayload mail.MessagePayload
		if err := json.Unmarshal(bodyBytes, &mailPayload); err == nil {
			// check the payload isn't empty
			if isValidMail(mailPayload) {
				storeMutex.Lock()
				emailStore = append(emailStore, mailPayload)
				storeMutex.Unlock()

				fmt.Printf("Email Received: %+v\n", mailPayload)
				c.JSON(http.StatusOK, gin.H{"status": "processed", "type": "mail"})
				return
			}
		}

		/*
		   var otherPayload OtherType
		   if err := json.Unmarshal(bodyBytes, &otherPayload); err == nil && otherPayload.ID != 0 {
		       // Handle other type
		       c.JSON(http.StatusOK, gin.H{"status": "processed", "type": "other"})
		       return
		   }
		*/

		fmt.Println("Received untyped webhook payload:", string(bodyBytes))
		c.JSON(http.StatusOK, gin.H{
			"status":    "received",
			"processed": true,
			"type":      "raw",
		})
	})

	r.Run(":" + port)
}

func isValidMail(m mail.MessagePayload) bool {
	return m.Data.From != ""
}
