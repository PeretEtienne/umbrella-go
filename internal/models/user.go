package models

import (
	"umbrella/internal/models/partial"
)

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Email          string `json:"email" gorm:"uniqueIndex;not null"`
	FirstName      string `json:"firstname" gorm:"size:255;not null"`
	LastName       string `json:"lastname" gorm:"size:255;not null"`
	IsCra          bool   `json:"isCra" gorm:"default:false"`
	HashedPassword string `json:"hashedPasswor" gorm:"size:255;not null"`
	RefreshToken   string `json:"refreshToken" gorm:"size:255"`
	IsActive       bool   `json:"isActive" gorm:"default:true"`
	partial.TimestampTracking
}
