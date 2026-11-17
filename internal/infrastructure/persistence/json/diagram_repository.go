package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

type diagramRepository struct {
	filePath string
}

// NewDiagramRepository creates a new JSON-backed DiagramRepository.
func NewDiagramRepository(dataDir string) repository.DiagramRepository {
	return &diagramRepository{
		filePath: filepath.Join(dataDir, "diagrams.json"),
	}
}

func (r *diagramRepository) GetModelingExperience() (*entity.ModelingExperience, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read diagrams file at %s: %w", r.filePath, err)
	}

	var experience entity.ModelingExperience
	if err := json.Unmarshal(data, &experience); err != nil {
		return nil, fmt.Errorf("failed to parse diagrams json: %w", err)
	}

	return &experience, nil
}

func (r *diagramRepository) GetDiagrams() ([]entity.Diagram, error) {
	exp, err := r.GetModelingExperience()
	if err != nil {
		return nil, err
	}
	return exp.Diagrams, nil
}

func (r *diagramRepository) GetDiagramByID(id string) (*entity.Diagram, error) {
	diagrams, err := r.GetDiagrams()
	if err != nil {
		return nil, err
	}

	for _, d := range diagrams {
		if d.ID == id {
			return &d, nil
		}
	}

	return nil, fmt.Errorf("diagram with ID '%s' not found", id)
}
