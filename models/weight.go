package models

type Weight_log struct {
	ID         uint  `json:"id" gorm:"primary_key"`
	Weight_log int64 `json:"Weight_log"`
	User_id    uint  `json:"User_id"`
}
