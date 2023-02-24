package model

type User struct {
	FirstName  string `json:"FirstName"`
	SecondName string `json:"SecondName"`
	Age        int    `json:"Age"`
	Email      string `json:"Email"`
}
