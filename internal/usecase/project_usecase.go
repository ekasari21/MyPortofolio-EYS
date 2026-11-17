package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// ProjectUseCase handles business logic for projects.
type ProjectUseCase interface {
	GetProjects() ([]entity.Project, error)
	GetFeaturedProjects() ([]entity.Project, error)
	GetProjectByID(id int) (*entity.Project, error)
}

type projectUseCase struct {
	repo repository.ProjectRepository
}

// NewProjectUseCase creates a new ProjectUseCase instance.
func NewProjectUseCase(repo repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{repo: repo}
}

func (u *projectUseCase) GetProjects() ([]entity.Project, error) {
	return u.repo.GetProjects()
}

func (u *projectUseCase) GetFeaturedProjects() ([]entity.Project, error) {
	return u.repo.GetFeaturedProjects()
}

func (u *projectUseCase) GetProjectByID(id int) (*entity.Project, error) {
	return u.repo.GetProjectByID(id)
}
