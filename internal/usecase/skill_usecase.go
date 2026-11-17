package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// SkillUseCase handles business logic for skills.
type SkillUseCase interface {
	GetSkills() ([]entity.SkillCategory, error)
}

type skillUseCase struct {
	repo repository.SkillRepository
}

// NewSkillUseCase creates a new SkillUseCase instance.
func NewSkillUseCase(repo repository.SkillRepository) SkillUseCase {
	return &skillUseCase{repo: repo}
}

func (u *skillUseCase) GetSkills() ([]entity.SkillCategory, error) {
	return u.repo.GetSkills()
}
