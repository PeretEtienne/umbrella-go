package services

import (
	"errors"
	"umbrella/internal/models"
	"umbrella/internal/repositories"
	"umbrella/internal/utils"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(password, user.HashedPassword) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
