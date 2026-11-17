package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// ExperienceUseCase handles business logic for work experiences.
type ExperienceUseCase interface {
	GetExperiences() ([]entity.Experience, error)
}

type experienceUseCase struct {
	repo repository.ExperienceRepository
}

// NewExperienceUseCase creates a new ExperienceUseCase instance.
func NewExperienceUseCase(repo repository.ExperienceRepository) ExperienceUseCase {
	return &experienceUseCase{repo: repo}
}

func (u *experienceUseCase) GetExperiences() ([]entity.Experience, error) {
	return u.repo.GetExperiences()
}
