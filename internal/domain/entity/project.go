package entity

// Project represents a portfolio case study or technical project.
type Project struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Category       string   `json:"category"`
	Role           string   `json:"role,omitempty"`
	Duration       string   `json:"duration,omitempty"`
	Description    string   `json:"description"`
	Problem        string   `json:"problem,omitempty"`
	Solution       string   `json:"solution,omitempty"`
	KeyFeatures    []string `json:"key_features,omitempty"`
	Impact         string   `json:"impact,omitempty"`
	Deliverables   []string `json:"deliverables,omitempty"`
	Technologies   []string `json:"technologies"`
	Image          string   `json:"image"`
	DemoURL        string   `json:"demo_url,omitempty"`
	RepositoryURL  string   `json:"repository_url,omitempty"`
	Featured       bool     `json:"featured,omitempty"`
}
