package main

import (
	_ "github.com/lib/pq"
	"solution/pkg/db"
	"solution/pkg/handler"
	nats_chanel "solution/pkg/nats-chanel"
)

func main() {
	db.RefreshDb()
	nats_chanel.Message()
	handler.Handle()
}
