package repository

import "portfolio/internal/domain/entity"

// ProjectRepository defines methods to access portfolio projects.
type ProjectRepository interface {
	GetProjects() ([]entity.Project, error)
	GetFeaturedProjects() ([]entity.Project, error)
	GetProjectByID(id int) (*entity.Project, error)
}
