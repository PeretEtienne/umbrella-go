package dto

type CreateMedicalDoctorDTO struct {
	Institution string `json:"institution" binding:"required,min=1,max=255"`
	Address     string `json:"address" binding:"required,min=1,max=255"`
	ZipCode     string `json:"zipCode" binding:"required,min=1,max=255"`
	City        string `json:"city" binding:"required,min=1,max=255"`
	Country     string `json:"country" binding:"required,min=1,max=255"`
}

type GetMedicalDoctorsPayloadDTO struct {
	Skip uint   `form:"skip"`
	Take uint   `form:"take"`
	Q    string `form:"q"`
}

type GetMedicalDoctorsResponseDTO struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"userId"`
	Email       string `json:"email"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	IsCra       bool   `json:"isCra"`
	Institution string `json:"institution"`
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
