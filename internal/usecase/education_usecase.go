package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"errors"
)

type EducationUsecase interface {
	GetAllByUserID(userID uint) ([]models.Education, error)
	Create(education *models.Education) error
	Update(userID, ID uint, education *models.Education) error
	Delete(userID, ID uint) error
}

type educationUsecase struct {
	repo repository.EducationRepository
}

func NewEducationUsecase(r repository.EducationRepository) EducationUsecase {
	return &educationUsecase{
		repo: r,
	}
}

func (u *educationUsecase) GetAllByUserID(userID uint) ([]models.Education, error) {
	return u.repo.GetAllByUserID(userID)
}

func (u *educationUsecase) Create(education *models.Education) error {

	if education.IsCurrent {
		education.EndDate = nil
	} else {
		if education.EndDate == nil {
			return errors.New("Tanggal selesai wajib diisi jika pendidikan sudah selesai")
		}

		if education.EndDate.Before(education.StartDate) {
			return errors.New("Tanggal selesai tidak boleh sebelum tanggal mulai")
		}
	}

	return u.repo.Create(education)
}

func (u *educationUsecase) Update(userID, ID uint, education *models.Education) error {

	if education.IsCurrent {
		education.EndDate = nil
	} else {
		if education.EndDate == nil {
			return errors.New("Tanggal selesai wajib diisi jika pendidikan sudah selesai")
		}

		if education.EndDate.Before(education.StartDate) {
			return errors.New("Tanggal selesai tidak boleh sebelum tanggal mulai")
		}
	}

	return u.repo.Update(userID, ID, education)
}

func (u *educationUsecase) Delete(userID, ID uint) error {
	return u.repo.Delete(userID, ID)
}
