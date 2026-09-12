package models

// Skill represents a technical skill and its proficiency level.
type Skill struct {
	Name        string `json:"name"`
	Proficiency int    `json:"proficiency,omitempty"` // percentage (e.g., 70, 80)
	Status      string `json:"status,omitempty"`      // e.g., "Ongoing", "Proficient"
}

// Project represents a portfolio project with detailed metadata.
type Project struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Summary      string   `json:"summary"`
	KeyFeatures  []string `json:"key_features"`
	Technologies []string `json:"technologies"`
	Goals        string   `json:"goals,omitempty"`
}

// BlogPost represents an article written by Sharat.
type BlogPost struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// SocialLinks holds profiles on social networks.
type SocialLinks struct {
	GitHub    string `json:"github"`
	LinkedIn  string `json:"linkedin"`
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
}

// Profile represents personal overview details.
type Profile struct {
	Name        string      `json:"name"`
	Headline    string      `json:"headline"`
	Tagline     string      `json:"tagline"`
	Quote       string      `json:"quote"`
	Bio         []string    `json:"bio"`
	Phone       string      `json:"phone"`
	Location    string      `json:"location"`
	SocialLinks SocialLinks `json:"social_links"`
}
