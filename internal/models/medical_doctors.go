package models

type MedicalDoctor struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Institution string `json:"institution" gorm:"size:255;not null"`
	Address     string `json:"address" gorm:"size:255;not null"`
	ZipCode     string `json:"zipCode" gorm:"size:255;not null"`
	City        string `json:"city" gorm:"size:255;not null"`
	Country     string `json:"country" gorm:"size:255;not null"`
	UserID      uint   `json:"userId" gorm:"not null"`
	User        User   `json:"user"`
}
