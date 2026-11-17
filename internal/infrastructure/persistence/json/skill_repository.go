package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type skillRepository struct {
	filePath string
}

// NewSkillRepository creates a new JSON-backed SkillRepository.
func NewSkillRepository(dataDir string) repository.SkillRepository {
	return &skillRepository{
		filePath: filepath.Join(dataDir, "skills.json"),
	}
}

func (r *skillRepository) GetSkills() ([]entity.SkillCategory, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read skills file at %s: %w", r.filePath, err)
	}

	var skills []entity.SkillCategory
	if err := json.Unmarshal(data, &skills); err != nil {
		return nil, fmt.Errorf("failed to parse skills json: %w", err)
	}

	return skills, nil
}
