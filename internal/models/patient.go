package models

//
// import (
// 	"umbrella/internal/constants"
// 	"umbrella/internal/models/partial"
// )
//
// type Patient struct {
// 	ID         uint                      `json:"id" gorm:"primaryKey"`
// 	ExternalID string                    `json:"externalId" gorm:"size:255;not null"`
// 	Initials   string                    `json:"initials" gorm:"size:255;not null"`
// 	BirthDate  string                    `json:"birthDate" gorm:"size:255;not null"`
// 	Gender     constants.Gender          `json:"gender" gorm:"type:enum('m', 'f'); not null; check(gender IN ('m', 'f'))"`
// 	Status     constants.Patientstatuses `json:"status" gorm:"type:enum('active', 'inactive'); not null; check(status IN ('active', 'inactive'))"`
// 	CenterID   uint                      `json:"centerId" gorm:"not null"`
// 	Center     Center                    `json:"center"`
// 	partial.TimestampTracking
// }
