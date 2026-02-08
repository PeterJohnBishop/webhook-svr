package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gin.SetMode(gin.ReleaseMode) // remove distracting logging
	r := gin.Default()

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
}
