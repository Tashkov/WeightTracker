package models

type User struct {
	Name   string `json:"Name"`
	Sex    string
	Age    int64 `json:"Age"`
	Height int64
}
