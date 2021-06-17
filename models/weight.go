package models

import "gorm.io/gorm"

type WeightLog struct {
	gorm.Model
	WeightLog int64 `json:"WeightLog"`
	UserID    uint  `json:"UserID"`
}
