package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"cv-builder-api/pkg"
	"errors"
)

type EducationUsecase interface {
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Education, pkg.PaginationMeta, error)
	Create(education *models.Education) error
	Update(userID, ID uint, education *models.Education) error
	Delete(userID, ID uint) error
	GetEducationByID(ID, userID uint) (*models.Education, error)
}

type educationUsecase struct {
	repo repository.EducationRepository
}

func NewEducationUsecase(r repository.EducationRepository) EducationUsecase {
	return &educationUsecase{
		repo: r,
	}
}

func (u *educationUsecase) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Education, pkg.PaginationMeta, error) {
	education, total, err := u.repo.GetAllByUserID(userID, pagination)

	if err != nil {
		return nil, pkg.PaginationMeta{}, err
	}

	meta := pkg.PaginationMeta{
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Total:     total,
		TotalPage: pkg.CalculateTotalPages(total, pagination.Limit),
		Filter:    pagination.Search,
		Sort:      pagination.Sort,
	}

	return education, meta, nil
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

func (u *educationUsecase) GetEducationByID(ID, userID uint) (*models.Education, error) {
	education, err := u.repo.GetEducationByID(ID, userID)

	if err != nil {
		return nil, errors.New("Data pendidikan tidak ditemukan")
	}

	return education, nil
}
