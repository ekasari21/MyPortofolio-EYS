package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type educationRepository struct {
	filePath string
}

// educationDataWrapper matches the JSON schema for education.json which can contain both education and certifications.
type educationDataWrapper struct {
	Education      []entity.Education     `json:"education"`
	Certifications []entity.Certification `json:"certifications"`
}

// NewEducationRepository creates a new JSON-backed EducationRepository.
func NewEducationRepository(dataDir string) repository.EducationRepository {
	return &educationRepository{
		filePath: filepath.Join(dataDir, "education.json"),
	}
}

func (r *educationRepository) load() (*educationDataWrapper, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read education file at %s: %w", r.filePath, err)
	}

	// Try parsing as educationDataWrapper first
	var wrapper educationDataWrapper
	if err := json.Unmarshal(data, &wrapper); err == nil && (len(wrapper.Education) > 0 || len(wrapper.Certifications) > 0) {
		return &wrapper, nil
	}

	// Fallback: try parsing as direct array []entity.Education
	var list []entity.Education
	if err := json.Unmarshal(data, &list); err == nil {
		return &educationDataWrapper{Education: list}, nil
	}

	return nil, fmt.Errorf("failed to parse education json: invalid structure")
}

func (r *educationRepository) GetEducation() ([]entity.Education, error) {
	wrapper, err := r.load()
	if err != nil {
		return nil, err
	}
	return wrapper.Education, nil
}

func (r *educationRepository) GetCertifications() ([]entity.Certification, error) {
	wrapper, err := r.load()
	if err != nil {
		return nil, err
	}
	return wrapper.Certifications, nil
}
