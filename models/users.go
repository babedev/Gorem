package models

type User struct {
	Fullname 	string	`json: "fullname"`
	Firstname 	string	`json: "firstname"`
	Lastname 	string	`json: "lastname"`
}

func Users() []User {
	return []User {
		User {"John Doe", "John", "Doe"},
		User {"John Doe", "John", "Doe"}
	}
}