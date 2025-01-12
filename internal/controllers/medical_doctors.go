package controllers

import (
	"fmt"
	"net/http"
	"umbrella/internal/dto"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
)

type MedicalDoctorController struct {
	UserService          *services.UserService
	MedicalDoctorService *services.MedicalDoctorService
}

func NewMedicalDoctorController(
	userService *services.UserService,
	medicalDoctorService *services.MedicalDoctorService,
) *MedicalDoctorController {
	return &MedicalDoctorController{
		UserService:          userService,
		MedicalDoctorService: medicalDoctorService,
	}
}

func (c *MedicalDoctorController) GetMedicalDoctors(ctx *gin.Context) {
	payload := dto.GetMedicalDoctorsPayloadDTO{
		Skip: 0,
		Take: 10,
		Q:    "",
	}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		fmt.Println("Error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doctors, err := c.MedicalDoctorService.GetMedicalDoctors(
		int(payload.Skip), int(payload.Take), payload.Q,
	)
	if err != nil {
		fmt.Println("Error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var response []dto.GetMedicalDoctorsResponseDTO
	for _, doctor := range doctors {
		response = append(response, dto.GetMedicalDoctorsResponseDTO{
			ID:          doctor.ID,
			UserID:      doctor.UserID,
			FirstName:   doctor.User.FirstName,
			LastName:    doctor.User.LastName,
			Email:       doctor.User.Email,
			IsCra:       doctor.User.IsCra,
			Institution: doctor.Institution,
			Address:     doctor.Address,
			ZipCode:     doctor.ZipCode,
			City:        doctor.City,
			Country:     doctor.Country,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
