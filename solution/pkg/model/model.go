package model

import "sync"

type User struct {
	Id         int    `json:"Id"`
	FirstName  string `json:"FirstName"`
	SecondName string `json:"SecondName"`
	Age        int    `json:"Age"`
	Email      string `json:"Email"`
}

var AllData []sync.Map

func AddData(element User) {
	var counter sync.Map
	counter.Store("first_name", element.FirstName)
	counter.Store("second_name", element.SecondName)
	counter.Store("age", element.Age)
	counter.Store("email", element.Email)
	AllData = append(AllData, counter)
}
