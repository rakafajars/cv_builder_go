package repository

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/pkg"
	"errors"

	"gorm.io/gorm"
)

type WorkExperinceRepository interface {
	Create(workExperience *models.WorkExperience) error
	Update(userID, ID uint, workExperience *models.WorkExperience) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.WorkExperience, int, error)
	GetWorkExperienceByID(ID, userID uint) (*models.WorkExperience, error)
}

type workExperinceRepository struct {
	db *gorm.DB
}

func NewWorkExperienceRepository(db *gorm.DB) WorkExperinceRepository {
	return &workExperinceRepository{db}
}

func (r *workExperinceRepository) Create(workExperience *models.WorkExperience) error {
	return r.db.Create(workExperience).Error
}

func (r *workExperinceRepository) Update(userID, ID uint, workExperience *models.WorkExperience) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Updates(workExperience)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data pengalaman kerja tidak ditemukan")
	}

	return nil
}

func (r *workExperinceRepository) Delete(userID, ID uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Delete(&models.WorkExperience{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data pengalaman kerja tidak ditemukan")
	}

	return nil
}

func (r *workExperinceRepository) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.WorkExperience, int, error) {
	var workExperience []models.WorkExperience
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if pagination.Search != "" {
		searchTerm := "%" + pagination.Search + "%"
		query = query.Where("company_name ILIKE ? OR position ILIKE ?", searchTerm, searchTerm)
	}

	err := query.Model(&models.WorkExperience{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	sortOrder := "start_date DESC"
	if pagination.Sort != "" {
		sortOrder = pagination.Sort
	}
	err = query.Limit(pagination.Limit).Offset(offset).Order(sortOrder).Find(&workExperience).Error

	return workExperience, int(total), err
}

func (r *workExperinceRepository) GetWorkExperienceByID(ID, userID uint) (*models.WorkExperience, error) {
	var workExperience models.WorkExperience
	err := r.db.Where("id = ? AND user_id = ?", ID, userID).First(&workExperience).Error

	return &workExperience, err
}
