package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// DiagramUseCase handles business logic for system modeling diagrams.
type DiagramUseCase interface {
	GetModelingExperience() (*entity.ModelingExperience, error)
	GetDiagrams() ([]entity.Diagram, error)
	GetDiagramByID(id string) (*entity.Diagram, error)
}

type diagramUseCase struct {
	repo repository.DiagramRepository
}

// NewDiagramUseCase creates a new DiagramUseCase instance.
func NewDiagramUseCase(repo repository.DiagramRepository) DiagramUseCase {
	return &diagramUseCase{repo: repo}
}

func (u *diagramUseCase) GetModelingExperience() (*entity.ModelingExperience, error) {
	return u.repo.GetModelingExperience()
}

func (u *diagramUseCase) GetDiagrams() ([]entity.Diagram, error) {
	return u.repo.GetDiagrams()
}

func (u *diagramUseCase) GetDiagramByID(id string) (*entity.Diagram, error) {
	return u.repo.GetDiagramByID(id)
}
