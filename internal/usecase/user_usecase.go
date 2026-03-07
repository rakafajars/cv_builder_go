package usecase

import (
	"cv-builder-api/internal/models"
	"cv-builder-api/internal/repository"
	"cv-builder-api/pkg"

	"errors"
)

type UserUsecase interface {
	Login(email, password string) (string, error)
	Register(email, password string) (*models.User, error)
}

type userUsecase struct {
	repo      repository.UserRepository
	secretKey string
}

func NewUserUsecase(r repository.UserRepository, secret string) UserUsecase {
	return &userUsecase{
		repo:      r,
		secretKey: secret,
	}
}

func (u *userUsecase) Register(email, password string) (*models.User, error) {
	// cek email sudah terdaftar apa belum

	_, err := u.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("Email sudah terdaftar")
	}

	// hash password dengan bcrpty
	hashedPassword, err := pkg.HashPassword(password)

	if err != nil {
		return nil, err
	}

	// buat object baru
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = u.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(email, password string) (string, error) {
	// mencari user di database
	user, err := u.repo.FindByEmail(email)

	if err != nil {
		return "", errors.New("Email atau Password Salah!")
	}

	if !pkg.CheckPasswordHash(password, user.Password) {
		return "", errors.New("Email atau Password Salah!")
	}

	tokenString, err := pkg.GenerateToken(user.ID, u.secretKey)

	return tokenString, err
}
