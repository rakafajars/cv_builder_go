package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"errors"
)

type SkillUsecase interface {
	GetAllByUserID(userID uint) ([]models.Skills, error)
	Create(skill *models.Skills) error
	Update(userID, ID uint, skill *models.Skills) error
	Delete(userID, id uint) error
}

func isValidLevel(level string) bool {
	validLevels := map[string]bool{
		"Beginner":     true,
		"Intermediate": true,
		"Advanced":     true,
		"Expert":       true,
	}

	return validLevels[level]
}

type skillUsecase struct {
	repo repository.SkillsRepository
}

func NewSkillsUsecase(r repository.SkillsRepository) SkillUsecase {
	return &skillUsecase{
		repo: r,
	}
}

func (u *skillUsecase) GetAllByUserID(userID uint) ([]models.Skills, error) {
	return u.repo.GetAllByUserID(userID)
}

func (u *skillUsecase) Create(skill *models.Skills) error {
	if !isValidLevel(skill.Level) {
		return errors.New("Level skill tidak valid, gunakan Beginner, Intermediate, Advanced Or Expert")
	}

	return u.repo.Create(skill)
}

func (u *skillUsecase) Update(userID, ID uint, skill *models.Skills) error {
	if !isValidLevel(skill.Level) {
		return errors.New("Level skill tidak valid, gunakan Beginner, Intermediate, Advanced Or Expert")
	}
	return u.repo.Update(userID, ID, skill)
}

func (u *skillUsecase) Delete(userID, id uint) error {
	return u.repo.Delete(userID, id)
}
