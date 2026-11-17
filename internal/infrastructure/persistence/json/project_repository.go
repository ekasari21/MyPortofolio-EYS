package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type projectRepository struct {
	filePath string
}

// NewProjectRepository creates a new JSON-backed ProjectRepository.
func NewProjectRepository(dataDir string) repository.ProjectRepository {
	return &projectRepository{
		filePath: filepath.Join(dataDir, "projects.json"),
	}
}

func (r *projectRepository) GetProjects() ([]entity.Project, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read projects file at %s: %w", r.filePath, err)
	}

	var projects []entity.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, fmt.Errorf("failed to parse projects json: %w", err)
	}

	return projects, nil
}

func (r *projectRepository) GetFeaturedProjects() ([]entity.Project, error) {
	projects, err := r.GetProjects()
	if err != nil {
		return nil, err
	}

	var featured []entity.Project
	for _, p := range projects {
		if p.Featured {
			featured = append(featured, p)
		}
	}

	// If no projects are marked featured explicitly, return the first 3
	if len(featured) == 0 && len(projects) > 0 {
		limit := 3
		if len(projects) < limit {
			limit = len(projects)
		}
		return projects[:limit], nil
	}

	return featured, nil
}

func (r *projectRepository) GetProjectByID(id int) (*entity.Project, error) {
	projects, err := r.GetProjects()
	if err != nil {
		return nil, err
	}

	for _, p := range projects {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("project with ID %d not found", id)
}
