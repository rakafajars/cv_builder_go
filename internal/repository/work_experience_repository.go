package repository

import (
	"cv-builder-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type WorkExperinceRepository interface {
	Create(workExperience *models.WorkExperience) error
	Update(userID, ID uint, workExperience *models.WorkExperience) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint) ([]models.WorkExperience, error)
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

func (r *workExperinceRepository) GetAllByUserID(userID uint) ([]models.WorkExperience, error) {
	var workExperience []models.WorkExperience

	err := r.db.Where("user_id = ?", userID).Find(&workExperience).Error

	return workExperience, err

}
