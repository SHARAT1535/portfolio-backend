package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"portfolio-backend/config"
	"portfolio-backend/handlers"
	"portfolio-backend/middleware"
)

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()

	// Handlers
	contactHandler := handlers.NewContactHandler(cfg)
	authHandler := handlers.NewAuthHandler(cfg)

	// API Endpoints
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/api/health", handlers.HealthHandler)
	mux.HandleFunc("/api/profile", handlers.ProfileHandler)
	mux.HandleFunc("/api/skills", handlers.SkillsHandler)
	mux.HandleFunc("/api/projects", handlers.ProjectsHandler)
	mux.HandleFunc("/api/blogs", handlers.BlogsHandler)

	mux.Handle("/api/contact", contactHandler)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)
	mux.HandleFunc("/api/auth/check", authHandler.CheckSession)

	// Middleware pipeline: Logger -> CORS -> Router
	handler := middleware.RequestLogger(middleware.EnableCORS(cfg, mux))

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("==================================================")
		fmt.Println("  🚀 Sharat R Kubasad - Portfolio Go Backend")
		fmt.Println("==================================================")
		fmt.Printf("  Server listening on:  http://localhost:%s\n", cfg.Port)
		fmt.Printf("  Health endpoint:      http://localhost:%s/api/health\n", cfg.Port)
		fmt.Printf("  Profile endpoint:     http://localhost:%s/api/profile\n", cfg.Port)
		fmt.Printf("  Skills endpoint:      http://localhost:%s/api/skills\n", cfg.Port)
		fmt.Printf("  Projects endpoint:    http://localhost:%s/api/projects\n", cfg.Port)
		fmt.Printf("  Blogs endpoint:       http://localhost:%s/api/blogs\n", cfg.Port)
		fmt.Printf("  Contact endpoint:     POST http://localhost:%s/api/contact\n", cfg.Port)
		fmt.Println("==================================================")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	<-stop
	log.Println("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exited cleanly.")
}
