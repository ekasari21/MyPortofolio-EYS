package repository

import "portfolio/internal/domain/entity"

// DiagramRepository defines methods to access system modeling diagrams and artifacts.
type DiagramRepository interface {
	GetModelingExperience() (*entity.ModelingExperience, error)
	GetDiagrams() ([]entity.Diagram, error)
	GetDiagramByID(id string) (*entity.Diagram, error)
}
