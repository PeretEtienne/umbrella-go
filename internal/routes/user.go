package routes

import (
	"umbrella/internal/controllers"
	"umbrella/internal/repositories"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
}
