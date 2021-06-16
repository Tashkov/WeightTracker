package models

import "gorm.io/gorm"

type Weight_log struct {
	gorm.Model
	Weight_log int64 `json:"Weight_log"`
	User_id    uint  `json:"User_id"`
}
