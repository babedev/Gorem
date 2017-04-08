package models

type User struct {
	Fullname  string        `json:"fullname"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Age       uint8        	`json:"age"`
}
