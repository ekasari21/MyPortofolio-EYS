package entity

// Experience represents a professional work experience or internship.
type Experience struct {
	Company          string   `json:"company"`
	Position         string   `json:"position"`
	Year             string   `json:"year"`
	Location         string   `json:"location,omitempty"`
	Type             string   `json:"type,omitempty"`
	Description      string   `json:"description"`
	Responsibilities []string `json:"responsibilities,omitempty"`
	Achievements     []string `json:"achievements,omitempty"`
	Technologies     []string `json:"technologies,omitempty"`
}
