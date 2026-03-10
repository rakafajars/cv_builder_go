package repository

import (
	"cv-builder-api/internal/models"

	"gorm.io/gorm"
)

type CVRepository interface {
	GetFullCV(userID uint) (*models.User, error)
}

type cvRepository struct {
	db *gorm.DB
}

func NewCVRepository(db *gorm.DB) CVRepository {
	return &cvRepository{db}
}

func (r *cvRepository) GetFullCV(userID uint) (*models.User, error) {
	var user models.User

	err := r.db.
		Preload("Profile").
		Preload("Experiences").
		Preload("Education").
		Preload("Skills").
		Preload("Projects").
		Where("id = ?", userID).
		First(&user).Error

	return &user, err
}
