package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"portfolio-backend/config"
	"portfolio-backend/models"
)

// ContactHandler handles contact submissions via JSON or Form.
type ContactHandler struct {
	cfg *config.Config
}

func NewContactHandler(cfg *config.Config) *ContactHandler {
	return &ContactHandler{cfg: cfg}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.ContactRequest

	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondJSON(w, http.StatusBadRequest, models.ApiResponse{
				Success: false,
				Message: "Invalid JSON request body",
				Error:   err.Error(),
			})
			return
		}
	} else {
		// Fallback for form-encoded submissions
		if err := r.ParseForm(); err != nil {
			respondJSON(w, http.StatusBadRequest, models.ApiResponse{
				Success: false,
				Message: "Failed to parse form submission",
				Error:   err.Error(),
			})
			return
		}
		req.Name = r.FormValue("name")
		req.Email = r.FormValue("email")
		req.Message = r.FormValue("message")
	}

	if err := req.Validate(); err != nil {
		respondJSON(w, http.StatusBadRequest, models.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
		return
	}

	log.Printf("[Contact] New message received from: %s <%s>\n", req.Name, req.Email)

	// Send email via SMTP
	if err := h.sendEmail(&req); err != nil {
		log.Printf("[Contact] SMTP warning/error: %v\n", err)
		// If SMTP credentials aren't configured yet, inform the user gently
		if h.cfg.SMTPPass == "" {
			respondJSON(w, http.StatusOK, models.ApiResponse{
				Success: true,
				Message: "Message received successfully! (SMTP_PASS not configured, logged to console)",
			})
			return
		}
		respondJSON(w, http.StatusInternalServerError, models.ApiResponse{
			Success: false,
			Message: "Failed to deliver email. Please try reaching out directly via LinkedIn or email.",
			Error:   err.Error(),
		})
		return
	}

	respondJSON(w, http.StatusOK, models.ApiResponse{
		Success: true,
		Message: "Thank you for your message! I will get back to you soon.",
	})
}

func (h *ContactHandler) sendEmail(req *models.ContactRequest) error {
	if h.cfg.SMTPPass == "" {
		return fmt.Errorf("SMTP password not set in environment")
	}

	addr := fmt.Sprintf("%s:%s", h.cfg.SMTPHost, h.cfg.SMTPPort)
	auth := smtp.PlainAuth("", h.cfg.SMTPUser, h.cfg.SMTPPass, h.cfg.SMTPHost)

	from := h.cfg.SMTPUser
	to := h.cfg.ReceiverEmail

	subject := "New message from website"
	body := fmt.Sprintf("From: %s <%s>\r\nReply-To: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\nFrom: %s <%s>\n\n%s",
		req.Name, req.Email, req.Email, subject, req.Name, req.Email, req.Message)

	msg := []byte(fmt.Sprintf("To: %s\r\n%s", to, body))

	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
