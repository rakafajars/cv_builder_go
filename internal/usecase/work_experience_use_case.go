package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
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

func (u *workExperienceUsecase) Create(workExperience *models.WorkExperience) error {
	return u.repo.Create(workExperience)
}

func (u *workExperienceUsecase) Update(userID, ID uint, workExperience *models.WorkExperience) error {
	return u.repo.Update(userID, ID, workExperience)
}

func (u *workExperienceUsecase) Delete(userID, id uint) error {
	return u.repo.Delete(userID, id)
}
