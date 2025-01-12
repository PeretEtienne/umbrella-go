package controllers

import (
	"fmt"
	"net/http"
	"umbrella/internal/dto"
	"umbrella/internal/models"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService          *services.UserService
	MedicalDoctorService *services.MedicalDoctorService
}

func NewUserController(
	userService *services.UserService,
	medicalDoctorService *services.MedicalDoctorService,
) *UserController {
	return &UserController{
		UserService:          userService,
		MedicalDoctorService: medicalDoctorService,
	}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userDTO dto.CreateUserPayloadDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		fmt.Println("Error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		IsCra:     userDTO.IsCra,
	}

	createdUser, err := c.UserService.CreateUser(user, userDTO.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userDTO.Type == "medical_doctor" {
		medicalDoctor := models.MedicalDoctor{
			Institution: userDTO.Institution,
			Address:     userDTO.Address,
			ZipCode:     userDTO.ZipCode,
			City:        userDTO.City,
			Country:     userDTO.Country,
			UserID:      createdUser.ID,
		}

		_, err := c.MedicalDoctorService.CreateMedicalDoctor(medicalDoctor)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	response := dto.CreateUserResponseDTO{
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		IsCra:     createdUser.IsCra,
	}

	ctx.JSON(http.StatusCreated, response)
}
