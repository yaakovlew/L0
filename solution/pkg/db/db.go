package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"solution/pkg/model"
	"solution/pkg/nats-chanel/validation"
)

const (
	userName = "postgres"
	password = ""
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
		allElementsFromDb.Scan(&id, &user.FirstName, &user.SecondName, &user.Email)
		model.AllData.Store(id, user)
	}
}

func AddInDB(m []byte) bool {
	if validation.ValidationMessage(m) {
		var user model.User
		err := json.Unmarshal(m, &user)
		if err != nil {
			panic(err)
		}
		connStr := fmt.Sprintf("user = %s password = %s dbname = %s sslmode=disable", userName, password, dbName)
		db, err := sql.Open("postgres", connStr)
		_, err = db.Exec("INSERT INTO users (first_name, second_name, email) values ($1, $2, $3)", user.FirstName, user.SecondName, user.Email)
		if err != nil {
			panic(err)
		}
		RefreshDb()
		return true
	} else {
		return false
	}
}
