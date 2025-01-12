package routes

import (
	"umbrella/internal/controllers"
	"umbrella/internal/repositories"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMedicalDoctorRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	medicalDoctorRepo := repositories.NewMedicalDoctorRepository(db)
	medicalDoctorService := services.NewMedicalDoctorService(medicalDoctorRepo)

	medicalDoctorController := controllers.NewMedicalDoctorController(userService, medicalDoctorService)

	router.GET("/medical-doctors", medicalDoctorController.GetMedicalDoctors)
}
