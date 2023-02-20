package model

import "sync"

type User struct {
	Id         int    `json:"Id"`
	FirstName  string `json:"FirstName"`
	SecondName string `json:"SecondName"`
	Age        int    `json:"Age"`
	Email      string `json:"Email"`
}

var AllData sync.Map
