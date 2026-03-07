package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"errors"
)

type ProfileUsecase interface {
	GetProfile(UserID uint) (*models.Profile, error)
	UpsertProfile(UserID int, fullName, phone, address, photoUrl, summary string) (*models.Profile, error)
}

type profileUsecase struct {
	repo repository.ProfileRepository
}

func NewProfileUsecase(r repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{
		repo: r,
	}
}

func (u *profileUsecase) GetProfile(UserID uint) (*models.Profile, error) {
	profile, err := u.repo.GetProfileByUserId(UserID)
	if err != nil {
		return nil, errors.New("Profile tidak ditemukan")
	}

	return profile, nil
}

func (u *profileUsecase) UpsertProfile(UserID int, fullName, phone, address, photoUrl, summary string) (*models.Profile, error) {

	profile := &models.Profile{
		UserID:   uint(UserID),
		FullName: fullName,
		Phone:    phone,
		Address:  address,
		PhotoUrl: photoUrl,
		Summary:  summary,
	}

	err := u.repo.Upsert(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
