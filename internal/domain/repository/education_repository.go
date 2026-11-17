package repository

import "portfolio/internal/domain/entity"

// EducationRepository defines methods to access education data.
type EducationRepository interface {
	GetEducation() ([]entity.Education, error)
	GetCertifications() ([]entity.Certification, error)
}
