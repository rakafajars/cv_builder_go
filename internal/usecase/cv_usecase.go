package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
)

type CVUsecase interface {
	GetCVData(userID uint) (*models.CVResponse, error)
}

type cvUsecase struct {
	repo repository.CVRepository
}

func NewCVUsecase(r repository.CVRepository) CVUsecase {
	return &cvUsecase{
		repo: r,
	}
}

func (u *cvUsecase) GetCVData(userID uint) (*models.CVResponse, error) {
	user, err := u.repo.GetFullCV(userID)

	if err != nil {
		return nil, err
	}

	response := &models.CVResponse{
		Profile:         &user.Profile,
		WorkExperiences: user.Experiences,
		Educations:      user.Education,
		Skills:          user.Skills,
		Projects:        user.Projects,
	}

	return response, nil
}
