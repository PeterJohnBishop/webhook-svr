package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"webhook-svr/mail"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v3"
)

var (
	emailStore   []resend.ReceivedEmail
	payloadStore []string
	storeMutex   sync.Mutex
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	resendToken := os.Getenv("RESEND_API_KEY")
	if resendToken == "" {
		log.Fatalln("Resend API Key not set!")
	}
	resendClient := resend.NewClient("re_xxxxxxxxx")

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
				mail, success := mail.GetMail(resendClient, mailPayload.Data.EmailID)
				if !success {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to fetch email"})
				}
				storeMutex.Lock()
				emailStore = append(emailStore, mail)
				storeMutex.Unlock()

				fmt.Printf("Email ID %s received.\n", mailPayload.Data.EmailID)
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
