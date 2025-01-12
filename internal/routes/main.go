package routes

import (
	"umbrella/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Public routes
	RegisterAuthRoutes(router, db)

	// Private routes
	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		RegisterUserRoutes(api, db)
	}

	return router
}
