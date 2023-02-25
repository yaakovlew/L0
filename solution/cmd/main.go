package main

import (
	_ "github.com/lib/pq"
	"solution/pkg/db"
	"solution/pkg/handler"
	nats_chanel "solution/pkg/nats-chanel"
)

func main() {
	db.RefreshDb()
	go nats_chanel.Produce()
	go nats_chanel.Take()
	handler.Handle()
}
