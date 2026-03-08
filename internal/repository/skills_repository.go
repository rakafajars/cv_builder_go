package repository

import (
	"cv-builder-api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type SkillsRepository interface {
	Create(skill *models.Skills) error
	Update(userID, ID uint, skill *models.Skills) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint) ([]models.Skills, error)
}

type skillsRepository struct {
	db *gorm.DB
}

func NewSkillsRepository(db *gorm.DB) SkillsRepository {
	return &skillsRepository{db}
}

func (r *skillsRepository) Create(skill *models.Skills) error {
	return r.db.Create(skill).Error
}

func (r *skillsRepository) Update(userID, ID uint, skill *models.Skills) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Updates(skill)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Skill tidak ditemukan")
	}

	return nil
}

func (r *skillsRepository) Delete(userID, ID uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userID, ID).Delete(&models.Skills{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Data Skill tidak ditemukan")
	}

	return nil
}

func (r *skillsRepository) GetAllByUserID(userID uint) ([]models.Skills, error) {
	var skill []models.Skills

	err := r.db.Where("user_id = ?", userID).Find(&skill).Error

	return skill, err

}
