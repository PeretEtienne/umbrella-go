package dto

type CreateUserDTO struct {
	FirstName string `json:"firstName" binding:"required,min=1,max=10"`
	LastName  string `json:"lastName" binding:"required,min=1,max=10"`
	Password  string `json:"password" binding:"required,min=8,max=255"`
	Email     string `json:"email" binding:"required,email"`
	IsCra     bool   `json:"isCra" default:"false"`
	Type      string `json:"type" binding:"required,oneof=medical_doctor user"`
}

type CreateUserPayloadDTO struct {
	CreateUserDTO
	CreateMedicalDoctorDTO
}

type CreateUserResponseDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	IsCra     bool   `json:"isCra"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}
