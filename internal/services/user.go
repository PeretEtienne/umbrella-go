package services

import (
	"umbrella/internal/models"
	"umbrella/internal/repositories"
	"umbrella/internal/utils"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.UserRepo.FindUserByID(id)
}

func (s *UserService) CreateUser(user models.User, password string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.HashedPassword = hashedPassword

	return s.UserRepo.CreateUser(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.UserRepo.UpdateUser(user)
}
