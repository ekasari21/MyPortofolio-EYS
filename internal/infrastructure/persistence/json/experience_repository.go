package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type experienceRepository struct {
	filePath string
}

// NewExperienceRepository creates a new JSON-backed ExperienceRepository.
func NewExperienceRepository(dataDir string) repository.ExperienceRepository {
	return &experienceRepository{
		filePath: filepath.Join(dataDir, "experiences.json"),
	}
}

func (r *experienceRepository) GetExperiences() ([]entity.Experience, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read experiences file at %s: %w", r.filePath, err)
	}

	var experiences []entity.Experience
	if err := json.Unmarshal(data, &experiences); err != nil {
		return nil, fmt.Errorf("failed to parse experiences json: %w", err)
	}

	return experiences, nil
}
