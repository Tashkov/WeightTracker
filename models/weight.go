package models

type Weight_log struct {
	ID         uint  `json:"id" gorm:"primary_key"`
	UserID     uint  `json:"userid" gorm:"foreign_key"`
	Weight_log int64 `json:"Weight_log"`
	Userid     int64 `json:"Userid"`
}
