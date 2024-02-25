package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateUserRequest represents the request structure for creating a user
type CreateUserRequest struct {
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}

func CreateUserHandler(c *gin.Context) {
	var createUserRequest CreateUserRequest

	// Bind JSON request body to createUserRequest struct
	if err := c.BindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate fields manually
	var errors []string
	if createUserRequest.FirstName == "" {
		errors = append(errors, "First name is required")
	}
	if createUserRequest.LastName == "" {
		errors = append(errors, "Last name is required")
	}
	if createUserRequest.Username == "" {
		errors = append(errors, "Username is required")
	}
	if createUserRequest.Email == "" {
		errors = append(errors, "Email is required")
	} else if !strings.Contains(createUserRequest.Email, "@") {
		errors = append(errors, "Invalid email format")
	}
	if createUserRequest.Password == "" {
		errors = append(errors, "Password is required")
	} else if len(createUserRequest.Password) < 6 {
		errors = append(errors, "Password must be at least 6 characters long")
	}

	// If there are validation errors, return them
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	// Here you would typically process the request, such as creating a user in a database
	// For this example, we'll just echo back the received data
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    createUserRequest,
	})
}
