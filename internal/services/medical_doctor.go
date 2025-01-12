package services

import (
	"umbrella/internal/models"
	"umbrella/internal/repositories"
)

type MedicalDoctorService struct {
	MedicalDoctorRepo *repositories.MedicalDoctorRepository
}

func NewMedicalDoctorService(medicalDoctorRepo *repositories.MedicalDoctorRepository) *MedicalDoctorService {
	return &MedicalDoctorService{MedicalDoctorRepo: medicalDoctorRepo}
}

func (s *MedicalDoctorService) GetMedicalDoctors(
	skip int, take int, q string,
) ([]models.MedicalDoctor, error) {
	return s.MedicalDoctorRepo.FindMedicalDoctors(skip, take, q)
}

func (s *MedicalDoctorService) CreateMedicalDoctor(medicalDoctor models.MedicalDoctor) (*models.MedicalDoctor, error) {
	return s.MedicalDoctorRepo.CreateMedicalDoctor(medicalDoctor)
}

func (s *MedicalDoctorService) Update(mediclDoctor *models.MedicalDoctor) error {
	return s.MedicalDoctorRepo.UpdateMedicalDoctor(mediclDoctor)
}
