package models

type User struct {
	ID        uint8			`json:"id"`
	Fullname  string        `json:"fullname"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Age       uint8        	`json:"age"`
	Contact	  Contact		`json:"contact"`
}
