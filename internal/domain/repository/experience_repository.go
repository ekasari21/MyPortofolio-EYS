package repository

import "portfolio/internal/domain/entity"

// ExperienceRepository defines methods to access experience data.
type ExperienceRepository interface {
	GetExperiences() ([]entity.Experience, error)
}
