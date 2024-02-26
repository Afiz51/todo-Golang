package validation

import (
	"strings"

	"github.com/Afiz51/TodoGoRest/models"
)

func ValidateUserRequest(req models.CreateUserRequest) []string {
	var errors []string

	if req.FirstName == "" {
		errors = append(errors, "First name is required")
	}
	if req.LastName == "" {
		errors = append(errors, "Last name is required")
	}
	if req.Username == "" {
		errors = append(errors, "Username is required")
	}
	if req.Email == "" {
		errors = append(errors, "Email is required")
	} else if !strings.Contains(req.Email, "@") {
		errors = append(errors, "Invalid email format")
	}
	if req.Password == "" {
		errors = append(errors, "Password is required")
	} else if len(req.Password) < 6 {
		errors = append(errors, "Password must be at least 6 characters long")
	}

	return errors
}
