package repository

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/pkg"
	"errors"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *models.Projects) error
	Update(userID, ID uint, project *models.Projects) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Projects, int, error)
	GetProjectByID(ID, userID uint) (*models.Projects, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db}
}

func (r *projectRepository) Create(project *models.Projects) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) Update(userID, ID uint, project *models.Projects) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Updates(project)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Project tidak ditemukan")
	}

	return nil
}

func (r *projectRepository) Delete(userID, ID uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Delete(&models.Projects{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Project tidak ditemukan")
	}

	return nil
}

func (r *projectRepository) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Projects, int, error) {
	var projects []models.Projects
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if pagination.Search != "" {
		searchTerm := "%" + pagination.Search + "%"
		query = query.Where("title ILIKE ? OR tech_stack ILIKE ? OR description ILIKE ?", searchTerm, searchTerm, searchTerm)
	}

	err := query.Model(&models.Projects{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	sortOrder := "created_at DESC"
	if pagination.Sort != "" {
		sortOrder = pagination.Sort
	}

	err = query.Limit(pagination.Limit).Offset(offset).Order(sortOrder).Find(&projects).Error

	return projects, int(total), err
}

func (r *projectRepository) GetProjectByID(ID, userID uint) (*models.Projects, error) {
	var project models.Projects
	err := r.db.Where("id = ? AND user_id = ?", ID, userID).First(&project).Error

	return &project, err

}
