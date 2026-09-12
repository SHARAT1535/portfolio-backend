package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"portfolio-backend/config"
	"portfolio-backend/models"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()

	HealthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var res HealthResponse
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if res.Status != "healthy" {
		t.Fatalf("expected healthy, got %s", res.Status)
	}
}

func TestPortfolioEndpoints(t *testing.T) {
	tests := []struct {
		name    string
		handler http.HandlerFunc
		path    string
	}{
		{"Profile", ProfileHandler, "/api/profile"},
		{"Skills", SkillsHandler, "/api/skills"},
		{"Projects", ProjectsHandler, "/api/projects"},
		{"Blogs", BlogsHandler, "/api/blogs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()
			tt.handler(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("%s handler returned %d", tt.name, rr.Code)
			}
			if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("expected Content-Type application/json, got %s", ct)
			}
		})
	}
}

func TestContactValidation(t *testing.T) {
	cfg := &config.Config{
		SMTPPass: "", // Test mode without SMTP
	}
	handler := NewContactHandler(cfg)

	// Test invalid email
	badPayload := `{"name":"Alice","email":"not-an-email","message":"Hello"}`
	req := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewBufferString(badPayload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for bad email, got %d", rr.Code)
	}

	// Test valid payload
	validPayload := `{"name":"Bob","email":"bob@example.com","message":"Hi Sharat!"}`
	req2 := httptest.NewRequest(http.MethodPost, "/api/contact", bytes.NewBufferString(validPayload))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()

	handler.ServeHTTP(rr2, req2)
	if rr2.Code != http.StatusOK {
		t.Fatalf("expected 200 for valid contact payload, got %d", rr2.Code)
	}
}

func TestAuthFlow(t *testing.T) {
	cfg := &config.Config{
		AdminUser:     "legion",
		AdminPassword: "password123",
	}
	auth := NewAuthHandler(cfg)

	// Valid login
	loginBody := `{"username":"legion","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBufferString(loginBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	auth.Login(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 for login, got %d", rr.Code)
	}

	var res models.LoginResponse
	_ = json.NewDecoder(rr.Body).Decode(&res)
	if !res.Success || res.Token == "" {
		t.Fatalf("expected successful login with token, got %+v", res)
	}

	// Check session
	checkReq := httptest.NewRequest(http.MethodGet, "/api/auth/check", nil)
	checkReq.Header.Set("Authorization", "Bearer "+res.Token)
	checkRR := httptest.NewRecorder()
	auth.CheckSession(checkRR, checkReq)
	if checkRR.Code != http.StatusOK {
		t.Fatalf("expected 200 for session check, got %d", checkRR.Code)
	}

	// Bad login
	badLogin := `{"username":"wrong","password":"wrong"}`
	badReq := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBufferString(badLogin))
	badReq.Header.Set("Content-Type", "application/json")
	badRR := httptest.NewRecorder()
	auth.Login(badRR, badReq)
	if badRR.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for bad login, got %d", badRR.Code)
	}
}
