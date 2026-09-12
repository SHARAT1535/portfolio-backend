# Portfolio Go Backend

A lightweight, high-performance REST API built in Go (Golang) using standard libraries. Migrated from Python/Flask to power the personal portfolio with clean JSON endpoints, CORS support for React, and SMTP contact handling.

## Architecture

- **Language**: Go 1.22+
- **Zero External Dependencies**: Uses Go's standard library (`net/http`, `net/smtp`, `encoding/json`, `crypto/rand`).
- **Features**:
  - RESTful endpoints for Profile, Skills, Projects, Blogs.
  - Contact form processing (JSON and Form POST) with Gmail SMTP email delivery.
  - Built-in CORS middleware for modern frontend (React/Vercel).
  - Configurable via environment variables or `.env` file.
  - Graceful server shutdown on interrupt signals.
  - Health check endpoint for uptime monitoring and AWS deployment.

---

## Directory Structure

```
gobackend/
├── go.mod                     # Go module declaration
├── main.go                    # Entry point & graceful server
├── config/
│   └── config.go              # Config loader & .env parser
├── data/
│   └── portfolio_data.go      # Portfolio content extracted from original templates
├── handlers/
│   ├── auth.go                # Admin authentication & session management
│   ├── contact.go             # Contact form & SMTP email dispatcher
│   ├── health.go              # Health & uptime monitoring
│   └── portfolio.go           # Profile, skills, projects, and blog endpoints
├── middleware/
│   ├── cors.go                # CORS middleware for React / Vercel
│   └── logger.go              # Request logging middleware
├── models/
│   ├── contact.go             # Contact form models & validators
│   ├── portfolio.go           # Portfolio domain structs
│   └── user.go                # Auth credentials struct
├── .env.example               # Example environment variables
└── README.md
```

---

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | API service overview and index |
| `GET` | `/api/health` | Service health status & uptime |
| `GET` | `/api/profile` | Personal bio, tagline, quote, and social links |
| `GET` | `/api/skills` | Technical skills with proficiency percentages |
| `GET` | `/api/projects` | Detailed project showcase with features & tech stack |
| `GET` | `/api/blogs` | Published articles & Medium walkthroughs |
| `POST` | `/api/contact` | Submit contact inquiry (`name`, `email`, `message`) |
| `POST` | `/api/auth/login` | Admin login (`username`, `password`) |
| `POST` | `/api/auth/logout` | Invalidate admin session |
| `GET` | `/api/auth/check` | Verify session validity |

---

## Running Locally

1. Navigate to the backend directory:
   ```bash
   cd gobackend
   ```

2. (Optional) Set up your environment file:
   ```bash
   cp .env.example .env
   ```

3. Run the Go server:
   ```bash
   go run main.go
   ```

4. Verify in your browser or terminal:
   ```bash
   curl http://localhost:8080/api/health
   curl http://localhost:8080/api/profile
   ```

---

## Building for Production / AWS

To compile into a standalone, single binary:

```bash
# For macOS (local)
go build -o server main.go

# For Linux (AWS EC2 / AWS Lambda / Render / Docker)
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server main.go
```
