package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Sex        string `json:"Sex"`
	Age        int64  `json:"Age"`
	Height_cm  int64  `json:"Height_cm"`
	Weight     []Weight_log
}
