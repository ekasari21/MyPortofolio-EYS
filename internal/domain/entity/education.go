package entity

// Education represents an academic qualification or formal certification.
type Education struct {
	Institution  string   `json:"institution"`
	Degree       string   `json:"degree"`
	FieldOfStudy string   `json:"field_of_study"`
	Year         string   `json:"year"`
	GPA          string   `json:"gpa,omitempty"`
	Description  string   `json:"description,omitempty"`
	Highlights   []string `json:"highlights,omitempty"`
}

// Certification represents a professional credential or industry license.
type Certification struct {
	Name         string `json:"name"`
	Issuer       string `json:"issuer"`
	Year         string `json:"year"`
	CredentialID string `json:"credential_id,omitempty"`
	URL          string `json:"url,omitempty"`
}
