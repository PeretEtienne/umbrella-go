package routes

import (
	"umbrella/internal/controllers"
	"umbrella/internal/repositories"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService, userService)

	router.POST("/auth/login", authController.Login)
}
