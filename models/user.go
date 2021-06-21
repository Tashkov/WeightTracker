package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Sex        string      `json:"sex"`
	Age        int64       `json:"age"`
	Height     int64       `json:"height"`
	WeightLogs []WeightLog `gorm:"foreignKey:UserID"`
}
