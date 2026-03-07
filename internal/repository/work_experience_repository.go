package repository

import (
	"cv-builder-api/internal/models"

	"gorm.io/gorm"
)

type WorkExperinceRepository interface {
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
	return r.db.Where("user_id = ? AND id = ?", userID, ID).Updates(workExperience).Error
}

func (r *workExperinceRepository) Delete(userID, ID uint) error {
	return r.db.Where("user_id = ? AND id = ?", userID, ID).Delete(&models.WorkExperience{}).Error
}

func (r *workExperinceRepository) GetAllByUserID(userID uint) ([]models.WorkExperience, error) {
	var workExperience []models.WorkExperience

	err := r.db.Where("user_id = ?", userID).Find(&workExperience).Error

	return workExperience, err

}
