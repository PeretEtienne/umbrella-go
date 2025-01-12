package models

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email" gorm:"uniqueIndex;not null"`
	FirstName      string    `json:"firstname" gorm:"size:255;not null"`
	LastName       string    `json:"lastname" gorm:"size:255;not null"`
	IsCra          bool      `json:"isCra" gorm:"default:false"`
	HashedPassword string    `json:"hashedPassword" gorm:"size:255;not null"`
	RefreshToken   string    `json:"refreshToken" gorm:"size:255"`
	CreatedAt      time.Time `json:"createAt" gorm:"autoCreateTime"`
	IsActive       bool      `json:"isActive" gorm:"default:true"`
	UpdatedAt      time.Time `json:"updateAt" gorm:"autoUpdateTime"`
}
