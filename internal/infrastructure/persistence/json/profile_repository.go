package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type profileRepository struct {
	filePath string
}

// NewProfileRepository creates a new JSON-backed ProfileRepository.
func NewProfileRepository(dataDir string) repository.ProfileRepository {
	return &profileRepository{
		filePath: filepath.Join(dataDir, "profile.json"),
	}
}

func (r *profileRepository) GetProfile() (*entity.Profile, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read profile file at %s: %w", r.filePath, err)
	}

	var profile entity.Profile
	if err := json.Unmarshal(data, &profile); err != nil {
		return nil, fmt.Errorf("failed to parse profile json: %w", err)
	}

	return &profile, nil
}
