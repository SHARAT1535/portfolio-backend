package handlers

import (
	"encoding/json"
	"net/http"

	"portfolio-backend/data"
)

// ProfileHandler responds with personal profile details.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respondJSON(w, http.StatusOK, data.GetProfile())
}

// SkillsHandler responds with list of skills.
func SkillsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respondJSON(w, http.StatusOK, data.GetSkills())
}

// ProjectsHandler responds with list of portfolio projects.
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respondJSON(w, http.StatusOK, data.GetProjects())
}

// BlogsHandler responds with list of blog posts.
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respondJSON(w, http.StatusOK, data.GetBlogs())
}

// RootHandler responds with API info and available endpoints.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	info := map[string]interface{}{
		"title":       "Sharat R Kubasad - Portfolio API",
		"version":     "1.0.0",
		"environment": "production",
		"endpoints": map[string]string{
			"GET  /api/health":   "Service health status",
			"GET  /api/profile":  "Developer profile, bio and contact info",
			"GET  /api/skills":   "Technical skills & proficiencies",
			"GET  /api/projects": "Detailed project showcases",
			"GET  /api/blogs":    "Published articles & walkthroughs",
			"POST /api/contact":  "Submit contact inquiry (JSON or Form)",
			"POST /api/auth/login": "Admin authentication",
		},
	}
	respondJSON(w, http.StatusOK, info)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
