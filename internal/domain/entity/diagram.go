package entity

// Diagram represents a system design, process workflow, or architectural diagram artifact.
type Diagram struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Category     string   `json:"category"` // e.g. "BPMN 2.0", "UML Sequence", "ERD (3NF)", "Architecture", "UML Use Case"
	Project      string   `json:"project"`
	Type         string   `json:"type"`
	Description  string   `json:"description"`
	KeyElements  []string `json:"key_elements"`
	Tools        []string `json:"tools"`
	MermaidCode  string   `json:"mermaid_code"`
	NDASanitized bool     `json:"nda_sanitized"`
}

// ModelingExperience summarizes the practitioner's system analysis and modeling competencies.
type ModelingExperience struct {
	Overview      string    `json:"overview"`
	Methodologies []string  `json:"methodologies"`
	Tools         []string  `json:"tools"`
	Stats         []Stat    `json:"stats"`
	Diagrams      []Diagram `json:"diagrams"`
}
