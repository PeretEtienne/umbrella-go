package repositories

import (
	"umbrella/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindUserByID(id string) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user models.User) (*models.User, error) {
	if err := repo.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.DB.Save(user).Error
}
