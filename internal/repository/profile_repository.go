package repository

import (
	"cv-builder-api/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProfileRepository interface {
	Create(profile *models.Profile) error
	Upsert(profile *models.Profile) error
	GetProfileByUserId(UserID uint) (*models.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}

func (r *profileRepository) Create(profile *models.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) Upsert(profile *models.Profile) error {
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"full_name",
			"phone",
			"address",
			"photo_url",
			"summary",
			"updated_at",
		}),
	}).Create(profile).Error
}

func (r *profileRepository) GetProfileByUserId(UserID uint) (*models.Profile, error) {
	var profile models.Profile

	err := r.db.Where("user_id = ?", UserID).First(&profile).Error

	return &profile, err
}
