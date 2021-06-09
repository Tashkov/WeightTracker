package models

type User struct {
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Sex        string
	Age        int64 `json:"Age"`
	Height_cm  int64 `json:"Height"`
}
