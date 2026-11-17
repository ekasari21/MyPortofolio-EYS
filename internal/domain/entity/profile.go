package entity

// Profile represents the professional profile of the portfolio owner.
type Profile struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone,omitempty"`
	Location    string   `json:"location"`
	LinkedIn    string   `json:"linkedin"`
	GitHub      string   `json:"github"`
	Summary     string   `json:"summary"`
	Bio         string   `json:"bio,omitempty"`
	ResumeURL   string   `json:"resume_url,omitempty"`
	KeyRoles    []string `json:"key_roles,omitempty"`
	Highlights  []string `json:"highlights,omitempty"`
	Stats       []Stat   `json:"stats,omitempty"`
}

// Stat represents a quantitative highlight (e.g. Projects Delivered, On-time rate).
type Stat struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Icon  string `json:"icon"`
}
