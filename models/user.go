package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string      `json:"FirstName"`
	LastName   string      `json:"LastName"`
	Sex        string      `json:"Sex"`
	Age        int64       `json:"Age"`
	Height     int64       `json:"Height"`
	WeightLogs []WeightLog `gorm:"foreignKey:UserID"`
}
