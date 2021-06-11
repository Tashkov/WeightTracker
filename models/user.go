package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Sex        string `json:"Sex"`
	Age        int64  `json:"Age"`
	Height_cm  int64  `json:"Height_cm"`
}
