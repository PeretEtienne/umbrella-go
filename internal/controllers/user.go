package controllers

import (
	"net/http"
	"umbrella/internal/dto"
	"umbrella/internal/models"
	"umbrella/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
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
	var userDTO dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
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

	response := dto.CreateUserResponseDTO{
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		IsCra:     createdUser.IsCra,
	}

	ctx.JSON(http.StatusCreated, response)
}
