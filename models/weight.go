package models

import "gorm.io/gorm"

type WeightLog struct {
	gorm.Model
	WeightLog int64 `json:"weight_log"`
	UserID    uint  `json:"user_id"`
}
