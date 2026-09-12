package models

// LoginRequest represents credentials for authentication.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login outcome.
type LoginResponse struct {
	Success  bool   `json:"success"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message"`
	Token    string `json:"token,omitempty"`
}
