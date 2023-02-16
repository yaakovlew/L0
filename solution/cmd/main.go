package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	userName = "postgres"
	password = "Andrei21.rus"
	dbName   = "wb"
)

func main() {

	connStr := fmt.Sprintf("user = %s password = %s dbname = %s sslmode=disable", userName, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO work(name) VALUES ('underfined')")
	if err != nil {
		panic(err)
	}
}
