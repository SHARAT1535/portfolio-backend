package models

import (
	"errors"
	"strings"
)

// ContactRequest holds the incoming payload for the contact form.
type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Validate checks that mandatory fields are present and reasonable.
func (c *ContactRequest) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(c.Email)
	c.Message = strings.TrimSpace(c.Message)

	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" || !strings.Contains(c.Email, "@") {
		return errors.New("valid email is required")
	}
	if c.Message == "" {
		return errors.New("message is required")
	}
	return nil
}

// ApiResponse represents a standard JSON API response.
type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
