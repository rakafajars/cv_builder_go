package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"errors"
)

type WorkExperienceUsecase interface {
	GetAllByUserID(userID uint) ([]models.WorkExperience, error)
	Delete(userID, id uint) error
	Create(workExperience *models.WorkExperience) error
	Update(userID, ID uint, workExperience *models.WorkExperience) error
}

type workExperienceUsecase struct {
	repo repository.WorkExperinceRepository
}

func NewWorkExperienceUsecase(r repository.WorkExperinceRepository) WorkExperienceUsecase {
	return &workExperienceUsecase{
		repo: r,
	}
}

func (u *workExperienceUsecase) GetAllByUserID(userID uint) ([]models.WorkExperience, error) {
	return u.repo.GetAllByUserID(userID)
}

func (u *workExperienceUsecase) Create(experience *models.WorkExperience) error {
	// pengecekan isCurrent jika masih bekerja
	// 1. flow jika is_current = true (Masih Kerja)
	if experience.IsCurrent {
		experience.EndDate = nil
	} else {
		// ATURAN 2: Kalau is_current = false (Sudah resign)
		// Pastikan end_date tidak kosong
		if experience.EndDate == nil {
			return errors.New("Tanggal selesai wabjib diisi jika sudah tidak bekerja")
		}

		if experience.EndDate.Before(experience.StartDate) {
			return errors.New("tanggal selesai tidak boleh sebelum tanggal mulai")
		}
	}
	return u.repo.Create(experience)
}

func (u *workExperienceUsecase) Update(userID, ID uint, experience *models.WorkExperience) error {
	// pengecekan isCurrent jika masih bekerja
	// 1. flow jika is_current = true (Masih Kerja)
	if experience.IsCurrent {
		experience.EndDate = nil
	} else {
		// ATURAN 2: Kalau is_current = false (Sudah resign)
		// Pastikan end_date tidak kosong
		if experience.EndDate == nil {
			return errors.New("Tanggal selesai wabjib diisi jika sudah tidak bekerja")
		}

		if experience.EndDate.Before(experience.StartDate) {
			return errors.New("tanggal selesai tidak boleh sebelum tanggal mulai")
		}
	}
	return u.repo.Update(userID, ID, experience)
}

func (u *workExperienceUsecase) Delete(userID, id uint) error {
	return u.repo.Delete(userID, id)
}
