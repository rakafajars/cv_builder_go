package repository

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/pkg"
	"errors"

	"gorm.io/gorm"
)

type EducationRepository interface {
	Create(education *models.Education) error
	Update(userID, ID uint, education *models.Education) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Education, int, error)
	GetEducationByID(ID, userID uint) (*models.Education, error)
}

type educationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db}
}

func (r *educationRepository) Create(education *models.Education) error {
	return r.db.Create(education).Error
}

func (r *educationRepository) Update(userID, ID uint, education *models.Education) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Updates(education)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Education tidak ditemukan")
	}

	return nil
}

func (r *educationRepository) Delete(userID, ID uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Delete(&models.Education{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Education tidak ditemukan")
	}

	return nil
}

func (r *educationRepository) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Education, int, error) {
	var education []models.Education
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if pagination.Search != "" {
		searchTerm := "%" + pagination.Search + "%"
		query = query.Where("institution ILIKE ? OR degree ILIKE ? OR field_of_study ILIKE ?", searchTerm, searchTerm, searchTerm)
	}

	err := query.Model(&models.Education{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	sortOrder := "created_at DESC"
	if pagination.Sort != "" {
		sortOrder = pagination.Sort
	}

	err = query.Limit(pagination.Limit).Offset(offset).Order(sortOrder).Find(&education).Error

	return education, int(total), err
}

func (r *educationRepository) GetEducationByID(ID, userID uint) (*models.Education, error) {
	var education models.Education
	err := r.db.Where("id = ? AND user_id = ?", ID, userID).First(&education).Error

	return &education, err
}
