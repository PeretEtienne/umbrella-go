package repositories

import (
	"umbrella/internal/models"

	"gorm.io/gorm"
)

type MedicalDoctorRepository struct {
	DB *gorm.DB
}

func NewMedicalDoctorRepository(db *gorm.DB) *MedicalDoctorRepository {
	return &MedicalDoctorRepository{DB: db}
}

func (repo *MedicalDoctorRepository) FindMedicalDoctorByID(id string) (*models.MedicalDoctor, error) {
	var medicalDoctor models.MedicalDoctor
	if err := repo.DB.First(&medicalDoctor, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &medicalDoctor, nil
}

func (repo *MedicalDoctorRepository) FindMedicalDoctors(skip int, take int, q string) ([]models.MedicalDoctor, error) {
	var medicalDoctors []models.MedicalDoctor
	if err := repo.DB.
		Preload("User").
		Joins("left join users on medical_doctors.user_id = users.id").
		Where("users.first_name LIKE ?", "%"+q+"%").Or("last_name LIKE ?", "%"+q+"%").
		Offset(skip).Limit(take).
		Find(&medicalDoctors).Error; err != nil {
		return nil, err
	}
	return medicalDoctors, nil
}

func (repo *MedicalDoctorRepository) CreateMedicalDoctor(medicalDoctor models.MedicalDoctor) (*models.MedicalDoctor, error) {
	if err := repo.DB.Create(&medicalDoctor).Error; err != nil {
		return nil, err
	}
	return &medicalDoctor, nil
}

func (repo *MedicalDoctorRepository) UpdateMedicalDoctor(medicalDoctor *models.MedicalDoctor) error {
	return repo.DB.Save(medicalDoctor).Error
}
