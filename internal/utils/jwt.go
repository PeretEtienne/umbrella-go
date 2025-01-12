package utils

import (
	"time"
	"umbrella/config"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.AccessTokenExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.RefreshTokenExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}
