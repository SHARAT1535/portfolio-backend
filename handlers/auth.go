package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"portfolio-backend/config"
	"portfolio-backend/models"
)

type AuthHandler struct {
	cfg      *config.Config
	sessions map[string]string // token -> username
	mu       sync.RWMutex
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg:      cfg,
		sessions: make(map[string]string),
	}
}

// Login handles admin authentication.
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondJSON(w, http.StatusBadRequest, models.LoginResponse{
				Success: false,
				Message: "Invalid JSON format",
			})
			return
		}
	} else {
		_ = r.ParseForm()
		req.Username = r.FormValue("username")
		req.Password = r.FormValue("password")
	}

	// Validate against configured credentials (from Python app: 'legion' / 'password123')
	if req.Username != a.cfg.AdminUser || req.Password != a.cfg.AdminPassword {
		respondJSON(w, http.StatusUnauthorized, models.LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		})
		return
	}

	token := generateToken()
	a.mu.Lock()
	a.sessions[token] = req.Username
	a.mu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	respondJSON(w, http.StatusOK, models.LoginResponse{
		Success:  true,
		Username: req.Username,
		Message:  "Login successful",
		Token:    token,
	})
}

// Logout clears the user session.
func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token := a.extractToken(r)
	if token != "" {
		a.mu.Lock()
		delete(a.sessions, token)
		a.mu.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	respondJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}

// CheckSession verifies if current request has an active valid session.
func (a *AuthHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	token := a.extractToken(r)
	if token == "" {
		respondJSON(w, http.StatusUnauthorized, models.ApiResponse{
			Success: false,
			Message: "No active session found",
		})
		return
	}

	a.mu.RLock()
	username, exists := a.sessions[token]
	a.mu.RUnlock()

	if !exists {
		respondJSON(w, http.StatusUnauthorized, models.ApiResponse{
			Success: false,
			Message: "Session expired or invalid",
		})
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"username": username,
		"message":  "Session active",
	})
}

func (a *AuthHandler) extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	cookie, err := r.Cookie("session_token")
	if err == nil {
		return cookie.Value
	}
	return ""
}

func generateToken() string {
	b := make([]byte, 24)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
