package repository

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/pkg"
	"errors"

	"gorm.io/gorm"
)

type SkillsRepository interface {
	Create(skill *models.Skills) error
	Update(userID, ID uint, skill *models.Skills) error
	Delete(userID, ID uint) error
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Skills, int, error)
	GetSkillByID(ID, userID uint) (*models.Skills, error)
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

func (r *skillsRepository) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Skills, int, error) {
	var skill []models.Skills
	var total int64

	query := r.db.Where("user_id = ?", userID)

	if pagination.Search != "" {
		searchTerm := "%" + pagination.Search + "%"
		query = query.Where("name ILIKE ? OR category ILIKE ? ", searchTerm, searchTerm)
	}

	err := query.Model(&models.Skills{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	sortOrder := "created_at DESC"
	if pagination.Sort != "" {
		sortOrder = pagination.Sort
	}

	err = query.Limit(pagination.Limit).Offset(offset).Order(sortOrder).Find(&skill).Error

	return skill, int(total), err

}

func (r *skillsRepository) GetSkillByID(ID, userID uint) (*models.Skills, error) {
	var skill models.Skills
	err := r.db.Where("id = ? AND user_id = ?", ID, userID).First(&skill).Error

	return &skill, err
}
