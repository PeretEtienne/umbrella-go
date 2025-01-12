package models

import (
	"umbrella/internal/models/partial"
)

type Center struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Code    string `json:"code" gorm:"uniqueIndex;size:255;not null"`
	Address string `json:"address" gorm:"size:255;not null"`
	ZipCode string `json:"zipCode" gorm:"size:255;not null"`
	City    string `json:"city" gorm:"size:255;not null"`
	Country string `json:"country" gorm:"size:255;not null"`
	partial.TimestampTracking
}
