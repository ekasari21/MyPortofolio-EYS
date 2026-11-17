package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// EducationUseCase handles business logic for education and credentials.
type EducationUseCase interface {
	GetEducation() ([]entity.Education, error)
	GetCertifications() ([]entity.Certification, error)
}

type educationUseCase struct {
	repo repository.EducationRepository
}

// NewEducationUseCase creates a new EducationUseCase instance.
func NewEducationUseCase(repo repository.EducationRepository) EducationUseCase {
	return &educationUseCase{repo: repo}
}

func (u *educationUseCase) GetEducation() ([]entity.Education, error) {
	return u.repo.GetEducation()
}

func (u *educationUseCase) GetCertifications() ([]entity.Certification, error) {
	return u.repo.GetCertifications()
}
