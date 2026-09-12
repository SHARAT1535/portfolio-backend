package config

import (
	"bufio"
	"os"
	"strings"
)

// Config holds the application configuration loaded from environment variables.
type Config struct {
	Port           string
	SMTPHost       string
	SMTPPort       string
	SMTPUser       string
	SMTPPass       string
	ReceiverEmail  string
	AllowedOrigins string
	AdminUser      string
	AdminPassword  string
}

// LoadConfig initializes the configuration from environment variables and an optional .env file.
func LoadConfig() *Config {
	loadDotEnv(".env")

	port := getEnv("PORT", "8080")
	smtpHost := getEnv("SMTP_HOST", "smtp.gmail.com")
	smtpPort := getEnv("SMTP_PORT", "587")
	smtpUser := getEnv("SMTP_USER", "sharat050604@gmail.com")
	smtpPass := getEnv("SMTP_PASS", "")
	receiverEmail := getEnv("RECEIVER_EMAIL", "samph37@gmail.com")
	allowedOrigins := getEnv("ALLOWED_ORIGINS", "*")
	adminUser := getEnv("ADMIN_USER", "legion")
	adminPassword := getEnv("ADMIN_PASSWORD", "password123")

	return &Config{
		Port:           port,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
		SMTPUser:       smtpUser,
		SMTPPass:       smtpPass,
		ReceiverEmail:  receiverEmail,
		AllowedOrigins: allowedOrigins,
		AdminUser:      adminUser,
		AdminPassword:  adminPassword,
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && strings.TrimSpace(val) != "" {
		return strings.TrimSpace(val)
	}
	return fallback
}

// loadDotEnv parses a key=value .env file if it exists without requiring external packages.
func loadDotEnv(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			if _, exists := os.LookupEnv(key); !exists {
				_ = os.Setenv(key, val)
			}
		}
	}
}
