package models

type User struct {
	Fullname 	string	`json: "fullname"`
	Firstname 	string	`json: "firstname"`
	Lastname 	string	`json: "lastname"`
	Age			uint8	`json: "age"`		
}

func Users() []User {
	return []User {
		User {"John Doe", "John", "Doe", 15},
		User {"John Doe", "John", "Doe", 17},
	}
}