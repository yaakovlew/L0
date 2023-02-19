package main

import (
	_ "github.com/lib/pq"
	"solution/pkg/db"
)

func main() {
	db.RefreshDb()
}
