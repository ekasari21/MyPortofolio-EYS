package entity

// SkillCategory groups technical and professional skills by competency domain.
type SkillCategory struct {
	Category    string   `json:"category"`
	Icon        string   `json:"icon,omitempty"`
	Description string   `json:"description,omitempty"`
	Items       []string `json:"items"`
}
