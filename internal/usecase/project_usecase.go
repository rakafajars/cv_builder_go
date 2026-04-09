package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"cv-builder-api/pkg"
	"errors"
)

type ProjectUsecase interface {
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Projects, pkg.PaginationMeta, error)
	Create(project *models.Projects) error
	Update(userID, ID uint, project *models.Projects) error
	Delete(userID, ID uint) error
	GetProjectByID(ID, userID uint) (*models.Projects, error)
}

type projectUsecase struct {
	repo repository.ProjectRepository
}

func NewProjectUsecase(r repository.ProjectRepository) ProjectUsecase {
	return &projectUsecase{
		repo: r,
	}
}

func (u *projectUsecase) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Projects, pkg.PaginationMeta, error) {
	projects, total, err := u.repo.GetAllByUserID(userID, pagination)

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

	return projects, meta, nil
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

func (u *projectUsecase) GetProjectByID(ID, userID uint) (*models.Projects, error) {
	project, err := u.repo.GetProjectByID(ID, userID)

	if err != nil {
		return nil, errors.New("Project tidak ditemukan")
	}

	return project, nil
}
