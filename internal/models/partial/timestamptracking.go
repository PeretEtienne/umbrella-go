package partial

import "time"

type TimestampTracking struct {
	CreatedAt time.Time `json:"createAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updateAt" gorm:"autoUpdateTime"`
}
