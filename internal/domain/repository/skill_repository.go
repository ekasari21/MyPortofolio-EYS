package repository

import "portfolio/internal/domain/entity"

// SkillRepository defines methods to access skill categories.
type SkillRepository interface {
	GetSkills() ([]entity.SkillCategory, error)
}
