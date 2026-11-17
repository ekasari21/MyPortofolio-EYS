package usecase

import (
	"portfolio/internal/domain/entity"
	"portfolio/internal/domain/repository"
)

// ProfileUseCase handles business logic for the profile.
type ProfileUseCase interface {
	GetProfile() (*entity.Profile, error)
}

type profileUseCase struct {
	repo repository.ProfileRepository
}

// NewProfileUseCase creates a new ProfileUseCase instance.
func NewProfileUseCase(repo repository.ProfileRepository) ProfileUseCase {
	return &profileUseCase{repo: repo}
}

func (u *profileUseCase) GetProfile() (*entity.Profile, error) {
	return u.repo.GetProfile()
}
