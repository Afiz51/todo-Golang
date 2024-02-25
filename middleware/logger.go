package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	// Initialize logger
	logger := log.New(os.Stdout, "[GIN-Logger] ", log.LstdFlags)

	return func(c *gin.Context) {
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		//print raw data
		fmt.Println("Raw data:", string(rawBody))

		c.Request.Body = io.NopCloser(bytes.NewReader(rawBody))

		// Log request details
		logger.Printf("[%s] %s %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP())

		// Proceed with the request
		c.Next()
	}
}
