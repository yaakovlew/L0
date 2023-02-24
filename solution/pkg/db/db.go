package db

import (
	"database/sql"
	"fmt"
	"solution/pkg/model"
)

const (
	userName = "postgres"
	password = "Andrei21.rus"
	dbName   = "wb"
)

func RefreshDb() {
	connStr := fmt.Sprintf("user = %s password = %s dbname = %s sslmode=disable", userName, password, dbName)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	allElementsFromDb, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	for allElementsFromDb.Next() {
		var user model.User
		var id int
		allElementsFromDb.Scan(&id, &user.FirstName, &user.SecondName, &user.Age, &user.Email)
		model.AllData.Store(id, user)
	}
}
