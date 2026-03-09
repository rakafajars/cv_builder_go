package repository

import (
	"cv-builder-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *models.Projects) error
	Update(userID, ID uint, project *models.Projects) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint) ([]models.Projects, error)
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

func (r *projectRepository) GetAllByUserID(userID uint) ([]models.Projects, error) {
	var projects []models.Projects

	err := r.db.Where("user_id = ?", userID).Find(&projects).Error

	return projects, err
}
