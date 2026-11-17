package repository

import "portfolio/internal/domain/entity"

// ProfileRepository defines methods to access profile data.
type ProfileRepository interface {
	GetProfile() (*entity.Profile, error)
}
