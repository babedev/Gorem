package models

type Contact struct {
	Email        string        `json:"email"`
	FacebookLink string        `json:"facebookLink"`
	TwitterLink  string        `json:"twitterLink"`
}
