package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"cv-builder-api/pkg"
	"errors"
)

type SkillUsecase interface {
	GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Skills, pkg.PaginationMeta, error)
	Create(skill *models.Skills) error
	Update(userID, ID uint, skill *models.Skills) error
	Delete(userID, id uint) error
	GetSkillByID(ID, userID uint) (*models.Skills, error)
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

func (u *skillUsecase) GetAllByUserID(userID uint, pagination pkg.PaginationQuery) ([]models.Skills, pkg.PaginationMeta, error) {
	skill, total, err := u.repo.GetAllByUserID(userID, pagination)

	if err != nil {
		return nil, pkg.PaginationMeta{}, err
	}

	meta := pkg.PaginationMeta{
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Total:     total,
		TotalPage: pkg.CalculateTotalPages(total, pagination.Limit),
		Filter:    pagination.Search,
		Sort:      pagination.Sort,
	}

	return skill, meta, nil
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

func (u *skillUsecase) GetSkillByID(ID, userID uint) (*models.Skills, error) {
	skill, err := u.repo.GetSkillByID(ID, userID)

	if err != nil {
		return nil, errors.New("Data skill tidak ditemukan")
	}

	return skill, nil
}
