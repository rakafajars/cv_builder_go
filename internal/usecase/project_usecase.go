package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
)

type ProjectUsecase interface {
	GetAllByUserID(userID uint) ([]models.Projects, error)
	Create(project *models.Projects) error
	Update(userID, ID uint, project *models.Projects) error
	Delete(userID, ID uint) error
}

type projectUsecase struct {
	repo repository.ProjectRepository
}

func NewProjectUsecase(r repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{
		repo: r,
	}
}

func (u *projectUsecase) GetAllByUserID(userID uint) ([]models.Projects, error) {
	return u.repo.GetAllByUserID(userID)
}

func (u *projectUsecase) Create(project *models.Projects) error {
	return u.repo.Create(project)
}

func (u *projectUsecase) Update(userID, ID uint, project *models.Projects) error {
	return u.repo.Update(userID, ID, project)
}

func (u *projectUsecase) Delete(userID, ID uint) error {
	return u.repo.Delete(userID, ID)
}
