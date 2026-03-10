package repository

import (
	"cv-builder-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type EducationRepository interface {
	Create(education *models.Education) error
	Update(userID, ID uint, education *models.Education) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint) ([]models.Education, error)
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

func (r *educationRepository) GetAllByUserID(userID uint) ([]models.Education, error) {
	var education []models.Education

	err := r.db.Where("user_id = ?", userID).Find(&education).Error

	return education, err

}
