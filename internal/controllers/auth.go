package controllers

import (
	"net/http"
	"umbrella/internal/dto"
	"umbrella/internal/services"
	"umbrella/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

func NewAuthController(
	authService *services.AuthService, userService *services.UserService,
) *AuthController {
	return &AuthController{AuthService: authService, UserService: userService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.AuthService.Login(loginDTO.Email, loginDTO.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	user.RefreshToken = refreshToken
	if err := c.UserService.Update(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user with refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"type":         "bearer",
	})
}
