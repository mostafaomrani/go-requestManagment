package models

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
}
